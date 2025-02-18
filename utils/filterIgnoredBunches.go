package utils

import "github.com/gotombola/godraw/types"

func FilterIgnoredBunches(data types.Data) ([]types.Bunch, error) {
	if len(data.IgnoredBunches) == 0 {
		return data.Bunches, nil
	}

	filteredBunchs := make([]types.Bunch, 0)
	ignoredBunches := make(map[string]int)

	for _, bunch := range data.IgnoredBunches {
		ignoredBunches[bunch.Id] = bunch.Nb
	}

	for _, bunch := range data.Bunches {
		if ignoredNb, exists := ignoredBunches[bunch.Id]; exists {
			bunch.Nb -= ignoredNb
			if bunch.Nb > 0 {
				filteredBunchs = append(filteredBunchs, bunch)
			}
		} else {
			filteredBunchs = append(filteredBunchs, bunch)
		}
	}

	return filteredBunchs, nil
}
