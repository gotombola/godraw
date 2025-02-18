package types

import "strconv"

type Options struct {
	Features       []string `json:"features,omitempty"`
	StartTimestamp int      `json:"sts,omitempty"`
	EndTimestamp   int      `json:"ets,omitempty"`
}

func (options Options) HasFeature(feature string) bool {
	for _, o := range options.Features {
		if o == feature {
			return true
		}
	}
	return false
}
func (options Options) GetMaxWinAmountPerOwnerFeature() int {
	for i := 1; i <= 3; i++ {
		found := options.HasFeature("max_" + strconv.Itoa(i) + "_per_owner")
		if found {
			return i
		}
	}
	return 0
}
func (options Options) GetMaxWinAmountPerTagPerOwnerFeature() int {
	for i := 1; i <= 3; i++ {
		found := options.HasFeature("max_" + strconv.Itoa(i) + "_per_tag_per_owner")
		if found {
			return i
		}
	}
	return 0
}
