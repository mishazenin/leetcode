package main

// items[i] = [type, color, name]
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var counter int
	keys := map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}
	for _, item := range items {
		intKey := keys[ruleKey]
		if item[intKey] == ruleValue {
			counter++
		}
	}

	return counter
}
