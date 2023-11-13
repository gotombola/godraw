package godraw

import (
	"time"

	"github.com/genstackio/gouuid"
)

//goland:noinspection GoUnusedExportedFunction
func CreateDraw(data Data) (Draw, error) {
	switch data.Mode {
	case "lottery":
		return createLotteryDraw(data)
	case "raffle":
		return createRaffleDraw(data)
	case "default":
		return createRaffleDraw(data)
	default:
		return createRaffleDraw(data)
	}
}

func filteredIgnoredTickets(data Data) ([]Ticket, error) {
	if len(data.IgnoredTickets) == 0 {
		return data.Tickets, nil
	}
	var filteredTickets []Ticket

	ticketsIgnored := make(map[string]bool)

	for _, ticket := range data.IgnoredTickets {
		ticketsIgnored[ticket.Id] = true
	}

	for _, ticket := range data.Tickets {
		if !ticketsIgnored[ticket.Id] {
			filteredTickets = append(filteredTickets, ticket)
		}
	}

	return filteredTickets, nil
}

func filteredIgnoredBunches(data Data) ([]Bunch, error) {
	if len(data.IgnoredBunches) == 0 {
		return data.Bunches, nil
	}

	filteredBunchs := make([]Bunch, 0)
	ignoredBunches := make(map[string]int)

	for _, bunch := range data.IgnoredBunches {
		ignoredBunches[bunch.Id] = bunch.Nb
	}

	for _, bunch := range data.Bunches {
		if ignoredNb, exists := ignoredBunches[bunch.Id]; exists {
			bunch.Nb -= ignoredNb
			if bunch.Nb > 0 {
				filteredBunchs = append(filteredBunchs, bunch)
			}
		} else {
			filteredBunchs = append(filteredBunchs, bunch)
		}
	}

	return filteredBunchs, nil
}

func createRaffleDraw(data Data) (Draw, error) {
	if len(data.Tickets) <= 0 {
		return Draw{
			Id:        gouuid.V4(),
			CreatedAt: time.Now().Format(time.RFC3339),
			Winners:   make([]Winner, 0),
		}, nil
	}

	if len(data.Bunches) <= 0 {
		return Draw{
			Id:        gouuid.V4(),
			CreatedAt: time.Now().Format(time.RFC3339),
			Winners:   make([]Winner, 0),
		}, nil
	}

	data.Bunches, _ = filteredIgnoredBunches(data)
	data.Tickets, _ = filteredIgnoredTickets(data)

	if data.PartialDraw && len(data.Tickets) > data.PartialMaxWinners {
		data.Tickets = data.Tickets[0:data.PartialMaxWinners]
	}

	nbShuffles := 5

	for i := 0; i < nbShuffles; i++ {
		data.Tickets = Shuffle(data.Tickets)
	}

	return Draw{
		Id:        gouuid.V4(),
		CreatedAt: time.Now().Format(time.RFC3339),
		Winners:   computeWinners(data.Tickets, data.Bunches),
	}, nil
}

func createLotteryDraw(data Data) (Draw, error) {
	owners := make(map[string]bool)

	var filteredTickets []Ticket

	for _, ticket := range data.Tickets {
		if len(ticket.Owner) != 0 && owners[ticket.Owner] {
			continue
		}

		filteredTickets = append(filteredTickets, ticket)

		if len(ticket.Owner) != 0 {
			owners[ticket.Owner] = true
		}
	}

	data.Tickets = filteredTickets

	return createRaffleDraw(data)
}
