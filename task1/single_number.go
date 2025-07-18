package task1

func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result = result ^ num
	}
	return result
}
