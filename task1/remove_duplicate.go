package task1

func RemoveDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if j < len(nums) && nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
