package types

type Winner struct {
	Ticket      string   `json:"t"`
	TicketOwner string   `json:"to"`
	Bunch       string   `json:"b"`
	TicketData  string   `json:"td,omitempty"`
	BunchData   string   `json:"bd,omitempty"`
	BunchTags   []string `json:"bt,omitempty"`
	RankOffset  int      `json:"ro,omitempty"`
}
