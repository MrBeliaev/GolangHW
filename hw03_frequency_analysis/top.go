package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var splitRegex = regexp.MustCompile(`\s+`)

func Top10(str string) []string {
	if strings.TrimSpace(str) == "" {
		return nil
	}
	strArr := splitRegex.Split(str, -1)
	wordsCount := map[string]int{}
	var maxCount int
	for _, v := range strArr {
		wordsCount[v]++
		if maxCount < wordsCount[v] {
			maxCount = wordsCount[v]
		}
	}
	result := make([]string, 0)
	for i := maxCount; i > 0; i-- {
		arr := make([]string, 0)
		for word, count := range wordsCount {
			if i == count {
				arr = append(arr, word)
			}
		}
		sort.Strings(arr)
		result = append(result, arr...)
	}
	if len(result) >= 10 {
		return result[:10]
	}
	return result
}
