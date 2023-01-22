package main

func main() {

}

// [7,1,14,11]
func checkIfExist(arr []int) bool {
	square := make(map[int]struct{})
	half := make(map[float64]struct{})
	for _, val := range arr {
		sq := val * 2
		hf := float64(val) / 2
		_, found1 := square[sq]
		_, found2 := half[hf]
		if found1 || found2 {
			return true
		}
		square[val] = struct{}{}
		half[float64(val)] = struct{}{}
	}
	return false
}
