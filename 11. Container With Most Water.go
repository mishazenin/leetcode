package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	lines := getLines(height)
	lines = func() []Line {
		input := lines
		for i := 0; i < len(input)-1; i++ {
			for j := 0; j < len(input)-1-i; j++ {
				if input[j].position > input[j+1].position {
					input[j], input[j+1] = input[j+1], input[j]
				}
			}
		}
		return input
	}()
	var max = 1

	heighest := lines[len(lines)-1]
	for i := len(lines) - 2; i > 0; i-- {
		currentHeight := heighest.heigth
		if lines[i].heigth < heighest.heigth {
			currentHeight = lines[i].heigth
		}
		lenght := heighest.position - lines[i].position
		currentVolume := getVolume(currentHeight, lenght)
		if currentVolume > max {
			max = currentVolume
		}

	}
	return max
}

func getVolume(height, lengh int) int {
	return height * lengh
}

func getLines(height []int) []Line {
	lines := make([]Line, len(height))
	for i, elem := range height {
		lines[i] = Line{
			heigth:   elem,
			position: i,
		}
	}
	return lines
}

type Line struct {
	heigth   int
	position int
}
