package godraw

import (
	"errors"
	"time"

	"github.com/genstackio/gouuid"
)

//goland:noinspection GoUnusedExportedFunction
func CreateDraw(data Data) (Draw, error) {
	if len(data.Tickets) <= 0 {
		return Draw{}, errors.New("not enough tickets")
	}

	if len(data.Bunches) <= 0 {
		return Draw{}, errors.New("not enough bunches")
	}

	nbShuffles := 5

	for i := 0; i < nbShuffles; i++ {
		data.Tickets = Shuffle(data.Tickets)
	}

	doc := Draw{
		Winners:   computeWinners(data.Tickets, data.Bunches),
		Id:        gouuid.V4(),
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	return doc, nil
}
