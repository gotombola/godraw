package utils

import "github.com/gotombola/godraw/types"

func AssignOne(bunch types.Bunch, ticket types.Ticket,
	counter int, giftCounter int, bunchNumber int, ticketNumber int,
	quantity int) (int, int, int, int, int, types.Winner) {
	if -1 == quantity {
		quantity = bunch.Quantity
	}
	var winner types.Winner
	if 0 < quantity {
		winner = types.Winner{Ticket: ticket.Id, TicketOwner: ticket.Owner, Bunch: bunch.Id, BunchTags: bunch.Tags, TicketData: ticket.Data, BunchData: bunch.Data}
		if bunch.RankOffset > 0 {
			winner.RankOffset = bunch.RankOffset
		}
		if 1 < quantity {
			giftCounter, quantity = giftCounter+1, quantity-1
		} else {
			bunchNumber, quantity = bunchNumber+1, -1
		}
		counter++
	} else {
		ticketNumber, bunchNumber, giftCounter, quantity =
			ticketNumber-1, bunchNumber+1, giftCounter-1, -1
	}
	return counter, bunchNumber, giftCounter, quantity, ticketNumber, winner
}
