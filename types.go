package godraw

import "strconv"

type Bunch struct {
	Id   string   `json:"id,omitempty"`
	Data string   `json:"data,omitempty"`
	Nb   int      `json:"nb,omitempty"`
	Ro   int      `json:"ro,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

type Winner struct {
	T  string   `json:"t"`
	To string   `json:"to"`
	B  string   `json:"b"`
	Td string   `json:"td,omitempty"`
	Bd string   `json:"bd,omitempty"`
	Bt []string `json:"bt,omitempty"`
	Ro int      `json:"ro,omitempty"`
}

type Draw struct {
	Id        interface{} `json:"id"`
	CreatedAt string      `json:"createdAt"`
	Winners   []Winner    `json:"winners"`
}

type Data struct {
	Tickets           []Ticket `json:"tickets,omitempty"`
	Bunches           []Bunch  `json:"bunches,omitempty"`
	Mode              string   `json:"mode,omitempty"`
	Features          []string `json:"features,omitempty"`
	IgnoredTickets    []Ticket `json:"ignoredTickets,omitempty"`
	PartialDraw       bool     `json:"partialDraw,omitempty"`
	PartialMaxWinners int      `json:"partialMaxWinners,omitempty"`
	IgnoredBunches    []Bunch  `json:"ignoredBunches,omitempty"`
}

type Ticket struct {
	Data          string   `json:"data,omitempty"`
	Id            string   `json:"id,omitempty"`
	Owner         string   `json:"owner,omitempty"`
	ChosenBunches []string `json:"chosenBunches,omitempty"`
}

func (ticket Ticket) hasOwnerAlreadyWonMaxAmount(ownersWins map[string][]Winner, max int) bool {
	wins, found := ownersWins[ticket.Owner]
	if !found {
		return false
	}
	if len(wins) >= max {
		return true
	}
	return false
}

func (ticket Ticket) hasOwnerAlreadyWonMaxAmountPerTag(ownersWinsTags map[string]map[string]int, max int, tags []string) bool {
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

func (ticket Ticket) hasChosenBunch(bunch string) bool {
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

type ComputeOptions struct {
	Features []string `json:"features,omitempty"`
}

func (options ComputeOptions) hasFeature(feature string) bool {
	for _, o := range options.Features {
		if o == feature {
			return true
		}
	}
	return false
}
func (options ComputeOptions) getMaxWinAmountPerOwnerFeature() int {
	for i := 1; i <= 3; i++ {
		found := options.hasFeature("max_" + strconv.Itoa(i) + "_per_owner")
		if found {
			return i
		}
	}
	return 0
}
func (options ComputeOptions) getMaxWinAmountPerTagPerOwnerFeature() int {
	for i := 1; i <= 3; i++ {
		found := options.hasFeature("max_" + strconv.Itoa(i) + "_per_tag_per_owner")
		if found {
			return i
		}
	}
	return 0
}
