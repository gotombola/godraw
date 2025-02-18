package utils

import "github.com/gotombola/godraw/types"

func FilterIgnoredTicketsBasedOnSameOwner(data types.Data) ([]types.Ticket, error) {
	owners := make(map[string]bool)
	var filteredTickets []types.Ticket

	for _, ticket := range data.Tickets {
		if len(ticket.Owner) != 0 && owners[ticket.Owner] {
			continue
		}

		filteredTickets = append(filteredTickets, ticket)

		if len(ticket.Owner) != 0 {
			owners[ticket.Owner] = true
		}
	}

	return filteredTickets, nil
}
