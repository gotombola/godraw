package modes

import (
	"github.com/genstackio/gouuid"
	"github.com/gotombola/godraw/types"
	"github.com/gotombola/godraw/utils"
	"time"
)

func Raffle(data types.Data) (types.Draw, error) {
	startTimestamp := time.Now().Second()
	draw := types.Draw{
		Id:        gouuid.V4(),
		CreatedAt: time.Now().Format(time.RFC3339),
		Winners:   make([]types.Winner, 0),
		Stats: types.Stats{
			NbTickets:      len(data.Tickets),
			StartTimestamp: startTimestamp,
		},
		StepStats: make([]types.StepStats, 0),
	}

	data.Bunches, _ = utils.FilterIgnoredBunches(data)
	data.Tickets, _ = utils.FilterIgnoredTickets(data)

	if len(data.Tickets) <= 0 {
		return draw, nil
	}

	if len(data.Bunches) <= 0 {
		return draw, nil
	}

	nbDraws := 0
	for _, b := range data.Bunches {
		nbDraws += b.Quantity
	}
	draw.Stats.NbDraws = nbDraws

	participant := map[string]bool{}
	for _, t := range data.Tickets {
		_, exist := participant[t.Owner]
		if !exist {
			participant[t.Owner] = true
		}
	}
	draw.Stats.NbParticipants = len(participant)

	if data.PartialDraw && len(data.Tickets) > data.PartialMaxWinners {
		data.Tickets = data.Tickets[0:data.PartialMaxWinners]
	}

	data.Tickets = utils.ShuffleTickets(data.Tickets)
	winners, stepStats := utils.ComputeWinners(
		data.Tickets,
		data.Bunches,
		types.Options{
			Features:             data.Features,
			TicketStartTimestamp: data.TicketStartTimestamp,
			TicketEndTimestamp:   data.TicketEndTimestamp,
		},
	)

	draw.Winners = winners
	draw.StepStats = stepStats

	participantWins := map[string]int{}
	max := 0
	for _, w := range winners {
		_, exist := participantWins[w.TicketOwner]
		if !exist {
			participantWins[w.TicketOwner] = 1
			if max == 0 {
				max = 1
			}
			continue
		}
		participantWins[w.TicketOwner]++
		if participantWins[w.TicketOwner] > max {
			max = participantWins[w.TicketOwner]
		}
	}

	winDistribution := map[int]int{}
	for i := 1; i < max+1; i++ {
		winDistribution[i] = 0
		for _, nb := range participantWins {
			if i == nb {
				winDistribution[i] += 1
			}
		}
	}
	sumNb := 0
	sumPercent := 0.0
	for nb, val := range winDistribution {
		percent := float64(val) * 100.0 / float64(draw.Stats.NbParticipants)
		draw.Stats.NbWinnersExactly = append(draw.Stats.NbWinnersExactly, types.WinnersStats{
			NbBunches: nb,
			Value:     val,
		})
		draw.Stats.PercentWinnersExactly = append(draw.Stats.PercentWinnersExactly, types.WinnersPercentStats{
			NbBunches: nb,
			Value:     percent,
		})

		sumNb += nb
		sumPercent += percent

		draw.Stats.NbWinnersAtLeast = append(draw.Stats.NbWinnersAtLeast, types.WinnersStats{
			NbBunches: nb,
			Value:     sumNb,
		})
		draw.Stats.PercentWinnersAtLeast = append(draw.Stats.PercentWinnersAtLeast, types.WinnersPercentStats{
			NbBunches: nb,
			Value:     sumPercent,
		})
	}

	endTimestamp := time.Now().Second()
	draw.Stats.EndTimestamp = endTimestamp
	draw.Stats.Duration = endTimestamp - startTimestamp
	return draw, nil
}
