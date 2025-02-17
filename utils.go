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
		winner = Winner{T: ticket.Id, To: ticket.Owner, B: bunch.Id, Bt: bunch.Tags, Td: ticket.Data, Bd: bunch.Data}
		if bunch.Ro > 0 {
			winner.Ro = bunch.Ro
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

// ComputeWinners creates a new list of 'giftCounter' from random numbers from the dataset
func computeWinners(tickets []Ticket, bunches []Bunch, options ComputeOptions) []Winner {
	var winners []Winner
	ownersWins := map[string][]Winner{}
	ownersWinsTags := map[string]map[string]int{}
	giftCounter := len(bunches)
	n := len(tickets)
	ticketNumber := 0
	if n > 1 {
		ticketNumber = rand.Intn(n - 1)
	}
	var bunchNumber, counter = 0, 0
	quantity := -1
	var winner Winner
	maxAmount := options.getMaxWinAmountPerOwnerFeature()
	maxAmountPerTag := options.getMaxWinAmountPerTagPerOwnerFeature()
	for ; counter < giftCounter && counter < n; ticketNumber++ {
		if maxAmount > 0 && tickets[ticketNumber].hasOwnerAlreadyWonMaxAmount(ownersWins, maxAmount) {
			if ticketNumber+1 == n {
				ticketNumber = -1
			}
			continue
		}
		if maxAmountPerTag > 0 && tickets[ticketNumber].hasOwnerAlreadyWonMaxAmountPerTag(ownersWinsTags, maxAmount, bunches[bunchNumber].Tags) {
			if ticketNumber+1 == n {
				ticketNumber = -1
			}
			continue
		}
		if !tickets[ticketNumber].hasChosenBunch(bunches[bunchNumber].Id) {
			if ticketNumber+1 == n {
				ticketNumber = -1
			}
			continue
		}
		counter, bunchNumber, giftCounter, quantity, ticketNumber, winner =
			assignOneOccurenceOfTheBunchToTheTicket(bunches[bunchNumber],
				tickets[ticketNumber], counter, giftCounter, bunchNumber, ticketNumber, quantity)
		if winner.T != "" {
			winners = append(winners, winner)
			ownersWins[winner.To] = append(ownersWins[winner.To], winner)
			if ownersWinsTags[winner.To] == nil {
				ownersWinsTags[winner.To] = map[string]int{}
			}
			for _, tag := range winner.Bt {
				ownersWinsTags[winner.To][tag]++
			}
		}
		if ticketNumber+1 == n {
			ticketNumber = -1
		}
	}

	return winners
}
