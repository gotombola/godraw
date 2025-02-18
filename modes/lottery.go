package modes

import (
	"github.com/gotombola/godraw/types"
	"github.com/gotombola/godraw/utils"
)

func Lottery(data types.Data) (types.Draw, error) {
	data.Tickets, _ = utils.FilterIgnoredTicketsBasedOnSameOwner(data)

	return Raffle(data)
}
