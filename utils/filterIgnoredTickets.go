package utils

import (
	"github.com/gotombola/godraw/types"
)

func FilterIgnoredTickets(data types.Data) ([]types.Ticket, error) {
	if len(data.IgnoredTickets) == 0 {
		return data.Tickets, nil
	}
	var filteredTickets []types.Ticket

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
