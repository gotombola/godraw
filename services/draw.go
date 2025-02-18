package services

import (
	"github.com/gotombola/godraw/modes"
	"github.com/gotombola/godraw/types"
)

var drawModes = map[string]types.ModeFn{
	"lottery": modes.Lottery,
	"raffle":  modes.Raffle,
	"default": modes.Raffle,
}

func CreateDraw(data types.Data) (types.Draw, error) {
	fn, ok := drawModes[data.Mode]

	if !ok {
		fn = drawModes["default"]
	}

	return fn(data)
}
