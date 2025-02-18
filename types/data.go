package types

type Data struct {
	Tickets           []Ticket `json:"tickets,omitempty"`
	Bunches           []Bunch  `json:"bunches,omitempty"`
	Mode              string   `json:"mode,omitempty"`
	Features          []string `json:"features,omitempty"`
	IgnoredTickets    []Ticket `json:"ignoredTickets,omitempty"`
	PartialDraw       bool     `json:"partialDraw,omitempty"`
	PartialMaxWinners int      `json:"partialMaxWinners,omitempty"`
	IgnoredBunches    []Bunch  `json:"ignoredBunches,omitempty"`
	StartTimestamp    int      `json:"sts,omitempty"`
	EndTimestamp      int      `json:"ets,omitempty"`
	Tags              []string `json:"tags,omitempty"`
}
