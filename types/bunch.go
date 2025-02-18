package types

import "encoding/json"

type Bunch struct {
	Id         string   `json:"id,omitempty"`
	Data       string   `json:"data,omitempty"`
	Quantity   int      `json:"nb,omitempty"`
	RankOffset int      `json:"ro,omitempty"`
	Tags       []string `json:"t,omitempty"`
}

func (b *Bunch) UnmarshalJSON(text []byte) error {
	type bunches Bunch

	bunch := bunches{
		Quantity: 1,
	}

	if err := json.Unmarshal(text, &bunch); err != nil {
		return err
	}

	*b = Bunch(bunch)

	return nil
}
