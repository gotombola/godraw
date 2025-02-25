package utils

import (
	"github.com/gotombola/godraw/types"
)

func FilterIgnoredTickets(data types.Data) ([]types.Ticket, error) {
	if len(data.IgnoredTickets) == 0 && data.TicketStartTimestamp == 0 && data.TicketEndTimestamp == 0 {
		return data.Tickets, nil
	}
	var filteredTickets []types.Ticket

	ticketsIgnored := make(map[string]bool)

	for _, ticket := range data.IgnoredTickets {
		ticketsIgnored[ticket.Id] = true
	}

	for _, ticket := range data.Tickets {
		if !ticketsIgnored[ticket.Id] && ticket.HasValidTimestamp(data.TicketStartTimestamp, data.TicketEndTimestamp) {
			filteredTickets = append(filteredTickets, ticket)
		}
	}

	return filteredTickets, nil
}
