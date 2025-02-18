package types

type Winner struct {
	T  string   `json:"t"`
	To string   `json:"to"`
	B  string   `json:"b"`
	Td string   `json:"td,omitempty"`
	Bd string   `json:"bd,omitempty"`
	Bt []string `json:"bt,omitempty"`
	Ro int      `json:"ro,omitempty"`
}
