package types

type Stats struct {
	NbTickets             int                   `json:"tickets"`
	NbDraws               int                   `json:"draws"`
	NbParticipants        int                   `json:"participants"`
	PercentWinnersAtLeast []WinnersPercentStats `json:"winnersAtLeastPercent"`
	NbWinnersAtLeast      []WinnersStats        `json:"winnersAtLeast"`
	PercentWinnersExactly []WinnersPercentStats `json:"winnersExactlyPercent"`
	NbWinnersExactly      []WinnersStats        `json:"winnersExactly"`
	StartTimestamp        int                   `json:"startTime"`
	EndTimestamp          int                   `json:"endTime"`
	Duration              int                   `json:"duration"`
	Steps                 []StepStats           `json:"steps"`
}

type StepStats struct {
	Index              int    `json:"i"`
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
	Index int `json:"i"`
	Value int `json:"v"`
}

type WinnersPercentStats struct {
	Index int     `json:"i"`
	Value float64 `json:"v"`
}
