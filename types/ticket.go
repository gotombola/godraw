package types

type Ticket struct {
	Data          string   `json:"data,omitempty"`
	Id            string   `json:"id,omitempty"`
	Owner         string   `json:"owner,omitempty"`
	ChosenBunches []string `json:"b,omitempty"`
	Timestamp     int      `json:"ts,omitempty"`
}

func (ticket Ticket) HasOwnerAlreadyWonMaxAmount(ownersWins map[string][]Winner, max int) bool {
	wins, found := ownersWins[ticket.Owner]
	if !found {
		return false
	}
	if len(wins) >= max {
		return true
	}
	return false
}

func (ticket Ticket) HasOwnerAlreadyWonMaxAmountPerTag(ownersWinsTags map[string]map[string]int, max int, tags []string) bool {
	owt, found := ownersWinsTags[ticket.Owner]
	if !found {
		return false
	}
	for _, tag := range tags {
		if owt[tag] > max {
			return true
		}
	}
	return false
}

func (ticket Ticket) HasChosenBunch(bunch string) bool {
	if ticket.ChosenBunches == nil || len(ticket.ChosenBunches) == 0 {
		return true
	}
	for _, cb := range ticket.ChosenBunches {
		if cb == "*" {
			return true
		}
		if cb == bunch {
			return true
		}
	}
	return false
}

func (ticket Ticket) HasValidTimestamp(start int, end int) bool {
	if start == 0 && end == 0 {
		return true
	}
	if ticket.Timestamp < start || (end != 0 && ticket.Timestamp > end) {
		return false
	}
	return true
}
