package main

// nums = [4,3,2,7,8,2,3,1]
// Тут внимание на условие задачи кол-во элементов равно длине массива т.о. элемент массива = его индекс+1
func findDisappearedNumbers(nums []int) []int {
	var stElems []int
	for _, i := range nums {
		idx := abs(i)
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	for i, val := range nums {
		if val > 0 {
			stElems = append(stElems, i+1)
		}
	}
	return stElems
}
func abs(elem int) int {
	if elem > 0 {
		return elem
	}
	return -elem
}
func main() {

}
