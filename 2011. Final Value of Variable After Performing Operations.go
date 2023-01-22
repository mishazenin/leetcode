package main

func finalValueAfterOperations(operations []string) int {
	var counter int
	for _, action := range operations {
		if action == "X++" || action == "++X" {
			counter++
		} else {
			counter--
		}
	}
	return counter
}
