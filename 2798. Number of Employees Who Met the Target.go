package main

// time O(n)
// space O(1)
func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var counter int
	for _, hour := range hours {
		if hour >= target {
			counter++
		}
	}
	return counter
}
