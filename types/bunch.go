package types

import "encoding/json"

type Bunch struct {
	Id   string   `json:"id,omitempty"`
	Data string   `json:"data,omitempty"`
	Nb   int      `json:"nb,omitempty"`
	Ro   int      `json:"ro,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

func (b *Bunch) UnmarshalJSON(text []byte) error {
	type bunches Bunch

	bunch := bunches{
		Nb: 1,
	}

	if err := json.Unmarshal(text, &bunch); err != nil {
		return err
	}

	*b = Bunch(bunch)

	return nil
}
