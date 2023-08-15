package godraw

import (
	"encoding/json"
	"math/rand"
)

func (b *Bunch) UnmarshalJSON(text []byte) error {
	type bunches Bunch

	bunch := bunches{
		Nb: 1,
	}

	if err := json.Unmarshal(text, &bunch); err != nil {
		return err
	}

	*b = Bunch(bunch)

	return nil
}

func Shuffle(items []Ticket) []Ticket {
	for i := len(items) - 1; i > 0; i-- {
		j := rand.Intn(i)
		if j != i {
			temp := items[i]
			items[i] = items[j]
			items[j] = temp
		}
	}
	return items
}

func assignOneOccurenceOfTheBunchToTheTicket(bunch Bunch, ticket Ticket,
	counter int, giftCounter int, bunchNumber int, ticketNumber int,
	quantity int) (int, int, int, int, int, Winner) {
	if -1 == quantity {
		quantity = bunch.Nb
	}
	var winner Winner
	if 0 < quantity {
		winner = Winner{T: ticket.Id, B: bunch.Id, Td: ticket.Data, Bd: bunch.Data}
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

// ComputeWinners creates a new list of 'giftCounter' from random numbers from the dataset
func computeWinners(tickets []Ticket, bunches []Bunch) []Winner {
	var winners []Winner
	giftCounter := len(bunches)
	var ticketNumber, bunchNumber, counter = rand.Intn(len(tickets) - 1), 0, 0
	quantity := -1
	var winner Winner
	for ; counter < giftCounter && counter < len(tickets); ticketNumber++ {
		counter, bunchNumber, giftCounter, quantity, ticketNumber, winner =
			assignOneOccurenceOfTheBunchToTheTicket(bunches[bunchNumber],
				tickets[ticketNumber], counter, giftCounter, bunchNumber, ticketNumber, quantity)
		if winner.T != "" {
			winners = append(winners, winner)
		}
		if ticketNumber+1 == len(tickets) {
			ticketNumber = -1
		}
	}

	return winners
}
