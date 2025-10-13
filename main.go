package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	array := []string{"abba", "baba", "bbaa", "cd", "cd"}
	res := removeAnagrams(array)
	fmt.Print(res)
}

func removeAnagrams(words []string) []string {
	var result []string
	var previousValue string

	for i := 0; i < len(words); i++ {
		currentValue := strings.Split(words[i], "")
		sort.Strings(currentValue)
		sortedValue := strings.Join(currentValue, "")

		if previousValue != sortedValue {
			result = append(result, words[i])
			previousValue = sortedValue
		}
	}
	return result
}
