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

	endTimestamp := time.Now().Second()
	draw.Stats.EndTimestamp = endTimestamp
	draw.Stats.Duration = endTimestamp - startTimestamp
	return draw, nil
}
