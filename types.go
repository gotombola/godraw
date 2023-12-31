package godraw

type Bunch struct {
	Id   string `json:"id,omitempty"`
	Data string `json:"data,omitempty"`
	Nb   int    `json:"nb,omitempty"`
	Ro   int    `json:"ro,omitempty"`
}

type Winner struct {
	T  string `json:"t"`
	B  string `json:"b"`
	Td string `json:"td,omitempty"`
	Bd string `json:"bd,omitempty"`
	Ro int    `json:"ro,omitempty"`
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
	IgnoredTickets    []Ticket `json:"ignoredTickets,omitempty"`
	PartialDraw       bool     `json:"partialDraw,omitempty"`
	PartialMaxWinners int      `json:"partialMaxWinners,omitempty"`
	IgnoredBunches    []Bunch  `json:"ignoredBunches,omitempty"`
}

type Ticket struct {
	Data  string `json:"data,omitempty"`
	Id    string `json:"id,omitempty"`
	Owner string `json:"owner,omitempty"`
}
