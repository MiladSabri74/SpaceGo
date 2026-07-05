package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	var output []string

	for _, word := range candidates {
		lowerWord := strings.ToLower(word)

		// Skip if same word or different length
		if lowerWord == subject || len(lowerWord) != len(subject) {
			continue
		}

		// Check if sorted characters match
		if sortString(subject) == sortString(lowerWord) {
			output = append(output, word)
		}
	}

	return output
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}
