package utils

import "maps"

func MapStringStructToSlice(input map[string]struct{}) []string {
	keys := make([]string, 0, len(input))
	for key := range maps.Keys(input) {
		keys = append(keys, key)
	}

	return keys
}
