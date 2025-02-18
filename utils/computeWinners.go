package utils

import (
	"github.com/gotombola/godraw/types"
	"math/rand"
)

func ComputeWinners(tickets []types.Ticket, bunches []types.Bunch, options types.Options) []types.Winner {
	var winners []types.Winner
	ownersWins := map[string][]types.Winner{}
	ownersWinsTags := map[string]map[string]int{}
	giftCounter := len(bunches)
	n := len(tickets)
	ticketNumber := 0
	if n > 1 {
		ticketNumber = rand.Intn(n - 1)
	}
	var bunchNumber, counter = 0, 0
	quantity := -1
	var winner types.Winner
	maxAmount := options.GetMaxWinAmountPerOwnerFeature()
	maxAmountPerTag := options.GetMaxWinAmountPerTagPerOwnerFeature()
	for ; counter < giftCounter && counter < n; ticketNumber++ {
		if maxAmount > 0 && tickets[ticketNumber].HasOwnerAlreadyWonMaxAmount(ownersWins, maxAmount) {
			if ticketNumber+1 == n {
				ticketNumber = -1
			}
			continue
		}
		if maxAmountPerTag > 0 && tickets[ticketNumber].HasOwnerAlreadyWonMaxAmountPerTag(ownersWinsTags, maxAmount, bunches[bunchNumber].Tags) {
			if ticketNumber+1 == n {
				ticketNumber = -1
			}
			continue
		}
		if !tickets[ticketNumber].HasChosenBunch(bunches[bunchNumber].Id) {
			if ticketNumber+1 == n {
				ticketNumber = -1
			}
			continue
		}
		counter, bunchNumber, giftCounter, quantity, ticketNumber, winner =
			AssignOne(bunches[bunchNumber],
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
