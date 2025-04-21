package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _,name := range names{
		runes := []rune(name)
		firstRune := runes[0]
		if nameCounts[firstRune] == nil {
			nameCounts[firstRune] = map[string]int{name:1}
			continue
		}
		nameCounts[firstRune][name]++
	}

	return nameCounts
}
