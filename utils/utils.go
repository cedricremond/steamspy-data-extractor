package utils

import (
	"strings"
)

func StringContainsCaseInsensitive(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func StringArrayToLower(strs []string) []string {
	newStrings := []string{}

	for _, str := range strs {
		newStrings = append(newStrings, strings.ToLower(str))
	}

	return newStrings
}
