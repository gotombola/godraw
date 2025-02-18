package utils

import (
	"math/rand"
)

func Shuffle[T interface{}](items []T) []T {
	for i := len(items) - 1; i > 0; i-- {
		j := rand.Intn(i)
		if j != i {
			temp := items[i]
			items[i] = items[j]
			items[j] = temp
		}
	}
	return items
}
