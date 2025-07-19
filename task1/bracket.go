package task1

func IsValid(s string) bool {
	var stk [10010]uint8
	var match = make(map[uint8]uint8)
	match[']'] = '['
	match['}'] = '{'
	match[')'] = '('
	tt := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if tt > 0 && stk[tt] == match[c] {
			tt--
		} else {
			tt++
			stk[tt] = c
		}
	}
	if tt == 0 {
		return true
	}
	return false
}
