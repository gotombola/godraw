package types

type Stats struct {
	NbTickets      int `json:"nb_tickets"`
	NbDraws        int `json:"nb_draws"`
	StartTimestamp int `json:"start_timestamp"`
	EndTimestamp   int `json:"end_timestamp"`
	Duration       int `json:"duration"`
}

type StepStats struct {
	Bunch              Bunch  `json:"b"`
	NbTickets          int    `json:"nbt"`
	NbBunches          int    `json:"nbb"`
	NbTicketsAfterDraw int    `json:"nbta"`
	NbBunchesAfterDraw int    `json:"nbtb"`
	Ticket             Ticket `json:"t"`
	StartTimestamp     int    `json:"sts"`
	EndTimestamp       int    `json:"ets"`
	Duration           int    `json:"d"`
}
