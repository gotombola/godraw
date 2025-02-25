package modes

import (
	"github.com/genstackio/gouuid"
	"github.com/gotombola/godraw/types"
	"github.com/gotombola/godraw/utils"
	"time"
)

func Raffle(data types.Data) (types.Draw, error) {
	startTimestamp := time.Now().UnixMilli()
	draw := types.Draw{
		Id:        gouuid.V4(),
		CreatedAt: time.Now().Format(time.RFC3339),
		Winners:   make([]types.Winner, 0),
		Stats: types.Stats{
			NbTickets:      len(data.Tickets),
			StartTimestamp: startTimestamp,
			Steps:          make([]types.StepStats, 0),
		},
	}

	data.Bunches, _ = utils.FilterIgnoredBunches(data)
	data.Tickets, _ = utils.FilterIgnoredTickets(data)

	if len(data.Tickets) <= 0 {
		return draw, nil
	}

	if len(data.Bunches) <= 0 {
		return draw, nil
	}

	bunchOwners := make(map[string]types.BunchStats, 0)
	nbDraws := 0
	for _, b := range data.Bunches {
		bunchOwners[b.Id] = types.BunchStats{
			Bunch:                          b.Id,
			NbOwnersHavingExplicitlyChosen: 0,
		}
		nbDraws += b.Quantity
	}
	draw.Stats.NbDraws = nbDraws

	owners := map[string]int{}
	for _, t := range data.Tickets {
		_, exist := owners[t.Owner]
		if !exist {
			owners[t.Owner] = 1
		} else {
			owners[t.Owner]++
		}
		if len(t.ChosenBunches) == 0 {
			continue
		}
		for _, b := range t.ChosenBunches {
			bs := bunchOwners[b]
			bs.NbOwnersHavingExplicitlyChosen++
			bunchOwners[b] = bs
		}
	}
	draw.Stats.NbOwners = len(owners)
	bs := make([]types.BunchStats, 0)
	for _, bo := range bunchOwners {
		bs = append(bs, bo)
	}
	draw.Stats.BunchStats = bs

	if data.PartialDraw && len(data.Tickets) > data.PartialMaxWinners {
		data.Tickets = data.Tickets[0:data.PartialMaxWinners]
	}

	data.Tickets = utils.ShuffleTickets(data.Tickets)
	winners, stepStats := utils.ComputeWinners(
		data.Tickets,
		data.Bunches,
		owners,
		types.Options{
			Features: data.Features,
		},
	)

	draw.Winners = winners
	draw.Stats.Steps = stepStats

	ownerWins := map[string]int{}
	vMax := 0
	for _, w := range winners {
		_, exist := ownerWins[w.TicketOwner]
		if !exist {
			ownerWins[w.TicketOwner] = 1
			if vMax == 0 {
				vMax = 1
			}
			continue
		}
		ownerWins[w.TicketOwner]++
		if ownerWins[w.TicketOwner] > vMax {
			vMax = ownerWins[w.TicketOwner]
		}
	}

	winDistribution := map[int]int{}
	for i := 1; i < vMax+1; i++ {
		winDistribution[i] = 0
		for _, nb := range ownerWins {
			if i == nb {
				winDistribution[i] += 1
			}
		}
	}
	for nb, val := range winDistribution {
		draw.Stats.NbWinnersExactly = append(draw.Stats.NbWinnersExactly, types.WinnersStats{
			Index: nb,
			Value: val,
		})
		draw.Stats.PercentWinnersExactly = append(draw.Stats.PercentWinnersExactly, types.WinnersPercentStats{
			Index: nb,
			Value: float64(val) / float64(draw.Stats.NbOwners),
		})
	}
	n := len(winDistribution)
	cumulativeNb := make([]int, n)
	cumulativePercent := make([]float64, n)

	runningNb := 0
	runningPercent := 0.0
	for i := n; i > 0; i-- {
		runningNb += winDistribution[i]
		cumulativeNb[i-1] = runningNb

		runningPercent += float64(winDistribution[i]) / float64(draw.Stats.NbOwners)
		cumulativePercent[i-1] = runningPercent
	}
	for i := 1; i <= n; i++ {
		draw.Stats.NbWinnersAtLeast = append(draw.Stats.NbWinnersAtLeast, types.WinnersStats{
			Index: i,
			Value: cumulativeNb[i-1],
		})
		draw.Stats.PercentWinnersAtLeast = append(draw.Stats.PercentWinnersAtLeast, types.WinnersPercentStats{
			Index: i,
			Value: cumulativePercent[i-1],
		})
	}

	endTimestamp := time.Now().UnixMilli()
	draw.Stats.EndTimestamp = endTimestamp
	draw.Stats.Duration = endTimestamp - startTimestamp
	return draw, nil
}
