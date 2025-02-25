package types

type Stats struct {
	NbTickets             int                   `json:"tickets"`
	NbDraws               int                   `json:"draws"`
	NbParticipants        int                   `json:"participants"`
	PercentWinnersAtLeast []WinnersPercentStats `json:"winnersAtLeastPercent"`
	NbWinnersAtLeast      []WinnersStats        `json:"winnersAtLeast"`
	PercentWinnersExactly []WinnersPercentStats `json:"winnersExactlyPercent"`
	NbWinnersExactly      []WinnersStats        `json:"winnersExactly"`
	StartTimestamp        int64                 `json:"startTime"`
	EndTimestamp          int64                 `json:"endTime"`
	Duration              int64                 `json:"duration"`
	Steps                 []StepStats           `json:"steps"`
}

type StepStats struct {
	Index              int    `json:"i"`
	Bunch              Bunch  `json:"bunch"`
	NbTickets          int    `json:"ticketsBefore"`
	NbTicketsAfterDraw int    `json:"ticketsAfter"`
	NbBunches          int    `json:"bunchesBefore"`
	NbBunchesAfterDraw int    `json:"bunchesAfter"`
	Ticket             Ticket `json:"ticket"`
	StartTimestamp     int64  `json:"startTime"`
	EndTimestamp       int64  `json:"endTime"`
	Duration           int64  `json:"duration"`
}

type WinnersStats struct {
	Index int `json:"index"`
	Value int `json:"value"`
}

type WinnersPercentStats struct {
	Index int     `json:"index"`
	Value float64 `json:"value"`
}
