package main

func containsDuplicate(nums []int) bool {

	m := make(map[int]struct{})
	for _, i := range nums {
		_, found := m[i]
		if found {
			return true
		}
		m[i] = struct{}{}
	}
	return false
}

func main() {

}
