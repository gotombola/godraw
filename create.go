package godraw

import (
	"errors"
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

func filteredIgnoreTickets(data Data) ([]Ticket, error) {
	if len(data.IgnoreTickets) == 0 {
		return data.Tickets, nil
	}
	var filteredTickets []Ticket

	ticketsIgnored := make(map[string]bool)

	for _, ticket := range data.IgnoreTickets {
		ticketsIgnored[ticket.Id] = true
	}

	for _, ticket := range data.Tickets {
		if !ticketsIgnored[ticket.Id] {
			filteredTickets = append(filteredTickets, ticket)
		}
	}

	return filteredTickets, nil
}

func createRaffleDraw(data Data) (Draw, error) {
	if len(data.Tickets) <= 0 {
		return Draw{}, errors.New("not enough tickets")
	}

	if len(data.Bunches) <= 0 {
		return Draw{}, errors.New("not enough bunches")
	}

	data.Tickets, _ = filteredIgnoreTickets(data)

	if data.PartialDraw && len(data.Tickets) > data.PartialMaxWinner {
		data.Tickets = data.Tickets[0:data.PartialMaxWinner]
	}

	nbShuffles := 5

	for i := 0; i < nbShuffles; i++ {
		data.Tickets = Shuffle(data.Tickets)
	}

	doc := Draw{
		Winners:   computeWinners(data.Tickets, data.Bunches),
		Id:        gouuid.V4(),
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	return doc, nil
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
