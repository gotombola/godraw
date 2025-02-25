package types

import "encoding/json"

type Bunch struct {
	Id         string   `json:"id,omitempty"`
	Data       string   `json:"data,omitempty"`
	Quantity   int      `json:"nb,omitempty"`
	RankOffset int      `json:"ro,omitempty"`
	Timestamp  int64    `json:"ts,omitempty"`
	Tags       []string `json:"tags,omitempty"`
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

func (bunch Bunch) HasValidTimestamp(start int64, end int64) bool {
	if start == 0 && end == 0 {
		return true
	}
	if bunch.Timestamp < start || (end != 0 && bunch.Timestamp > end) {
		return false
	}
	return true
}
