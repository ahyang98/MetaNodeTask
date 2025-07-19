package task1

func PlusOne(digits []int) []int {
	over := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + over
		digits[i] = sum % 10
		over = sum / 10
	}
	if over > 0 {
		ret := make([]int, 1)
		ret[0] = over
		return append(ret, digits...)
	}
	return digits
}
