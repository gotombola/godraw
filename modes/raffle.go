package modes

import (
	"github.com/genstackio/gouuid"
	"github.com/gotombola/godraw/types"
	"github.com/gotombola/godraw/utils"
	"time"
)

func Raffle(data types.Data) (types.Draw, error) {
	draw := types.Draw{
		Id:        gouuid.V4(),
		CreatedAt: time.Now().Format(time.RFC3339),
		Winners:   make([]types.Winner, 0),
	}

	data.Bunches, _ = utils.FilterIgnoredBunches(data)
	data.Tickets, _ = utils.FilterIgnoredTickets(data)

	if len(data.Tickets) <= 0 {
		return draw, nil
	}

	if len(data.Bunches) <= 0 {
		return draw, nil
	}

	if data.PartialDraw && len(data.Tickets) > data.PartialMaxWinners {
		data.Tickets = data.Tickets[0:data.PartialMaxWinners]
	}

	data.Tickets = utils.ShuffleTickets(data.Tickets)
	draw.Winners = utils.ComputeWinners(
		data.Tickets,
		data.Bunches,
		types.Options{
			Features:       data.Features,
			StartTimestamp: data.StartTimestamp,
			EndTimestamp:   data.EndTimestamp,
		},
	)

	return draw, nil
}
