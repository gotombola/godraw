package types

type Stats struct {
	NbTickets             int                   `json:"nb_tickets"`
	NbDraws               int                   `json:"nb_draws"`
	NbParticipants        int                   `json:"nb_participants"`
	PercentWinnersAtLeast []WinnersPercentStats `json:"percent_winners_at_least"`
	NbWinnersAtLeast      []WinnersStats        `json:"nb_winners_at_least"`
	PercentWinnersExactly []WinnersPercentStats `json:"percent_winners_exactly"`
	NbWinnersExactly      []WinnersStats        `json:"nb_winners_exactly"`
	StartTimestamp        int                   `json:"start_timestamp"`
	EndTimestamp          int                   `json:"end_timestamp"`
	Duration              int                   `json:"duration"`
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

type WinnersStats struct {
	NbBunches int `json:"nbb"`
	Value     int `json:"v"`
}

type WinnersPercentStats struct {
	NbBunches int     `json:"nbb"`
	Value     float64 `json:"v"`
}
