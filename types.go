package godraw

type Bunch struct {
	Id   string
	Data string
	Nb   int
}

type Winner struct {
	T  string `json:"t"`
	B  string `json:"b"`
	Td string `json:"td,omitempty"`
	Bd string `json:"bd,omitempty"`
}

type Draw struct {
	Id        interface{} `json:"id"`
	CreatedAt string      `json:"createdAt"`
	Winners   []Winner    `json:"winners"`
}

type Data struct {
	Tickets          []Ticket `json:"tickets,omitempty"`
	Bunches          []Bunch  `json:"bunches,omitempty"`
	Mode             string   `json:"mode,omitempty"`
	IgnoreTickets    []Ticket
	PartialDraw      bool
	PartialMaxWinner int
}

type Ticket struct {
	Data  string `json:"data,omitempty"`
	Id    string `json:"id,omitempty"`
	Owner string `json:"owner,omitempty"`
}
