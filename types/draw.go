package types

type Draw struct {
	Id        interface{} `json:"id"`
	CreatedAt string      `json:"createdAt"`
	Winners   []Winner    `json:"winners"`
}
