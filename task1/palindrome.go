package task1

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	num := strconv.Itoa(x)

	j := len(num) - 1
	for i := 0; i < len(num); i++ {
		if num[i] != num[j] {
			return false
		}
		j--
		if j <= i {
			return true
		}
	}
	return false
}
