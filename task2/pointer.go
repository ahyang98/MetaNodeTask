package task2

import "fmt"

func setValue(p *int, v int) {
	*p = v
}

func TestSetValue() {
	a := 10
	setValue(&a, 11)
	println(a)
}

func setSliceValue(s []int) {
	for i, _ := range s {
		s[i] = s[i] * 2
	}
}

func TestSlice() {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	setSliceValue(s)
	fmt.Println(s)
}

func setArrayValue(a *[3]int) {
	for i, _ := range a {
		a[i] = a[i] * 2
	}
}

func TestArray() {
	a := [3]int{1, 2, 3}
	setArrayValue(&a)
	fmt.Println(a)
}
