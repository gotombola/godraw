package types

type Data struct {
	Tickets              []Ticket `json:"tickets,omitempty"`
	Bunches              []Bunch  `json:"bunches,omitempty"`
	Mode                 string   `json:"mode,omitempty"`
	Features             []string `json:"features,omitempty"`
	IgnoredTickets       []Ticket `json:"ignoredTickets,omitempty"`
	PartialDraw          bool     `json:"partialDraw,omitempty"`
	PartialMaxWinners    int      `json:"partialMaxWinners,omitempty"`
	IgnoredBunches       []Bunch  `json:"ignoredBunches,omitempty"`
	TicketStartTimestamp int64    `json:"ticketStartTime,omitempty"`
	TicketEndTimestamp   int64    `json:"ticketEndTime,omitempty"`
	BunchStartTimestamp  int64    `json:"bunchStartTime,omitempty"`
	BunchEndTimestamp    int64    `json:"bunchEndTime,omitempty"`
	Tags                 []string `json:"tags,omitempty"`
}
