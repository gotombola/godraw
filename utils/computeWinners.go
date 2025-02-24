package utils

import (
	"github.com/gotombola/godraw/types"
	"math/rand"
	"time"
)

func ComputeWinners(tickets []types.Ticket, bunches []types.Bunch, options types.Options) ([]types.Winner, []types.StepStats) {
	var winners []types.Winner
	var stepStats []types.StepStats
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
		step := types.StepStats{
			Bunch:          bunches[bunchNumber],
			NbTickets:      n - counter,
			NbBunches:      giftCounter - counter,
			Ticket:         tickets[ticketNumber],
			StartTimestamp: time.Now().Second(),
		}
		if (maxAmount > 0 && tickets[ticketNumber].HasOwnerAlreadyWonMaxAmount(ownersWins, maxAmount)) ||
			maxAmountPerTag > 0 && tickets[ticketNumber].HasOwnerAlreadyWonMaxAmountPerTag(ownersWinsTags, maxAmount, bunches[bunchNumber].Tags) ||
			!tickets[ticketNumber].HasChosenBunch(bunches[bunchNumber].Id) ||
			!tickets[ticketNumber].HasValidTimestamp(options.TicketStartTimestamp, options.TicketEndTimestamp) {
			if ticketNumber+1 == n {
				ticketNumber = -1
			}
			endTimestamp := time.Now().Second()
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
		if ticketNumber+1 == n {
			ticketNumber = -1
		}
		endTimestamp := time.Now().Second()
		step.EndTimestamp = endTimestamp
		step.Duration = endTimestamp - step.StartTimestamp
		step.NbTicketsAfterDraw = n - counter
		step.NbBunchesAfterDraw = giftCounter - counter
		stepStats = append(stepStats, step)

	}

	return winners, stepStats
}
