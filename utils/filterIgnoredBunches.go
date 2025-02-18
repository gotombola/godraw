package utils

import "github.com/gotombola/godraw/types"

func FilterIgnoredBunches(data types.Data) ([]types.Bunch, error) {
	tagsLength := len(data.Tags)
	if len(data.IgnoredBunches) == 0 && tagsLength == 0 {
		return data.Bunches, nil
	}

	filteredBunchs := make([]types.Bunch, 0)
	ignoredBunches := make(map[string]int)

	for _, bunch := range data.IgnoredBunches {
		ignoredBunches[bunch.Id] = bunch.Quantity
	}

	tagSet := make(map[string]struct{}, tagsLength)
	for _, tag := range data.Tags {
		tagSet[tag] = struct{}{}
	}

	for _, bunch := range data.Bunches {
		if ignoredNb, exists := ignoredBunches[bunch.Id]; exists {
			bunch.Quantity -= ignoredNb
			if bunch.Quantity <= 0 {
				continue
			}
		}
		if tagsLength == 0 {
			filteredBunchs = append(filteredBunchs, bunch)
			continue
		}
		for _, tag := range bunch.Tags {
			if _, found := tagSet[tag]; found {
				filteredBunchs = append(filteredBunchs, bunch)
				break
			}
		}
	}

	return filteredBunchs, nil
}
