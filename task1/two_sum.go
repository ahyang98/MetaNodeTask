package task1

func TwoSum(nums []int, target int) []int {
	idx := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		idx[nums[i]] = i
	}
	ret := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		if v, ok := idx[target-nums[i]]; ok && v != i {

			ret[0] = i
			ret[1] = v
			break
		}
	}
	return ret
}
