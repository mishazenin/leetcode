package main

import "fmt"

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
}
func sortPeople(names []string, heights []int) []string {
	for i := 0; i < len(heights)-1; i++ {
		for j := 0; j < len(heights)-i-1; j++ {
			if heights[j] < heights[j+1] {
				heights[j], heights[j+1] = heights[j+1], heights[j]
				names[j], names[j+1] = names[j+1], names[j]
			}

		}
	}
	return names
}
