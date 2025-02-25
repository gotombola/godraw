package utils

import (
	"github.com/gotombola/godraw/types"
	"math/rand"
	"time"
)

func ComputeWinners(tickets []types.Ticket, bunches []types.Bunch, owners map[string]int, options types.Options) ([]types.Winner, []types.StepStats) {
	var winners []types.Winner
	var stepStats []types.StepStats
	ownersWins := map[string][]types.Winner{}
	ownersWinsTags := map[string]map[string]int{}
	giftCounter := len(bunches)
	n := len(tickets)
	ticketNumber := 0
	ignoredOwnersTicketsCounter := 0
	if n > 1 {
		ticketNumber = rand.Intn(n - 1)
	}
	var bunchNumber, counter = 0, 0
	quantity := -1
	var winner types.Winner
	maxAmount := options.GetMaxWinAmountPerOwnerFeature()
	maxAmountPerTag := options.GetMaxWinAmountPerTagPerOwnerFeature()
	for ; counter < giftCounter && counter < n; ticketNumber++ {
		step := types.StepStats{
			Index:          giftCounter - counter,
			Bunch:          bunches[bunchNumber].Id,
			NbTickets:      n - counter - ignoredOwnersTicketsCounter,
			NbBunches:      giftCounter - counter,
			Ticket:         tickets[ticketNumber].Id,
			StartTimestamp: time.Now().UnixMilli(),
		}
		if (maxAmount > 0 && tickets[ticketNumber].HasOwnerAlreadyWonMaxAmount(ownersWins, maxAmount)) ||
			maxAmountPerTag > 0 && tickets[ticketNumber].HasOwnerAlreadyWonMaxAmountPerTag(ownersWinsTags, maxAmount, bunches[bunchNumber].Tags) ||
			!tickets[ticketNumber].HasChosenBunch(bunches[bunchNumber].Id) {
			if ticketNumber+1 == n {
				ticketNumber = -1
			}
			endTimestamp := time.Now().UnixMilli()
			step.EndTimestamp = endTimestamp
			step.Duration = endTimestamp - step.StartTimestamp
			step.NbTicketsAfterDraw = n - counter
			step.NbBunchesAfterDraw = giftCounter - counter
			stepStats = append(stepStats, step)
			continue
		}
		counter, bunchNumber, giftCounter, quantity, ticketNumber, winner =
			AssignOne(bunches[bunchNumber],
				tickets[ticketNumber], counter, giftCounter, bunchNumber, ticketNumber, quantity)
		if winner.Ticket != "" {
			winners = append(winners, winner)
			ownersWins[winner.TicketOwner] = append(ownersWins[winner.TicketOwner], winner)
			if ownersWinsTags[winner.TicketOwner] == nil {
				ownersWinsTags[winner.TicketOwner] = map[string]int{}
			}
			for _, tag := range winner.BunchTags {
				ownersWinsTags[winner.TicketOwner][tag]++
			}
		}
		if maxAmount > 0 && tickets[ticketNumber].HasOwnerAlreadyWonMaxAmount(ownersWins, maxAmount) {
			ignoredOwnersTicketsCounter += owners[winner.TicketOwner] - 1
		}
		if ticketNumber+1 == n {
			ticketNumber = -1
		}
		endTimestamp := time.Now().UnixMilli()
		step.EndTimestamp = endTimestamp
		step.Duration = endTimestamp - step.StartTimestamp
		//@warning: ignoredOwnersTicketsCounter does not take the max_per_tag_per_owner limit into account
		step.NbTicketsAfterDraw = n - counter - ignoredOwnersTicketsCounter
		step.NbBunchesAfterDraw = giftCounter - counter
		stepStats = append(stepStats, step)

	}

	return winners, stepStats
}
