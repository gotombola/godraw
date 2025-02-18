package utils

import "github.com/gotombola/godraw/types"

func ShuffleTickets(tickets []types.Ticket) []types.Ticket {
	nbShuffles := 5

	for i := 0; i < nbShuffles; i++ {
		tickets = Shuffle(tickets)
	}

	return tickets
}
