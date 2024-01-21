package main

import "fmt"

func main() {
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20))

}

// 1,1,2,3,4
func maxIceCream(costs []int, coins int) int {
	var counter int
	for _, price := range bubbleSort(costs) {
		if coins-price >= 0 {
			counter++
			coins = coins - price
		} else {
			break
		}
	}
	return counter
}

// func maxIceCream(costs []int, coins int) int {
// 	counter := 0
// 	var maxNumbers int
// 	costs = bubbleSort(costs)
// 	restMoney := coins
// 	var next int
// 	for _, price := range costs {
// 		if next != 0 {
// 			price = next
// 		}
// 		// enough
// 		if restMoney-price >= 0 {
// 			counter++
// 			restMoney = restMoney - price
// 			if counter > maxNumbers {
// 				maxNumbers = counter
// 			}
// 		} else {
// 			counter = 0
// 			restMoney = coins
// 		}
// 	}
// 	return maxNumbers
// }

// func bubbleSort(arr []int) []int {
// 	for i := 0; i < len(arr)-1; i++ {
// 		for j := 0; j < len(arr)-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// 	return arr
// }
