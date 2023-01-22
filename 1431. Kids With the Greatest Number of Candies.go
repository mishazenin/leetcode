package main

// candies = [2,3,5,1,3], extraCandies = 3
func kidsWithCandies(candies []int, extraCandies int) []bool {

	var res []bool
	var max int
	for _, num := range candies {
		max = maxCands(max, num)
	}

	for _, val := range candies {
		if val+extraCandies > max {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}

func maxCands(a, b int) int {
	if a > b {
		return a
	}
	return b
}
