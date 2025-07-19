package task1

import "sort"

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := make([][]int, 0)
	idx := 0
	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		r := ret[idx][1]
		nl := intervals[i][0]
		nr := intervals[i][1]
		if nl <= r {
			if nr > r {
				ret[idx][1] = nr
			}
		} else {
			idx++
			newRange := make([]int, 2)
			newRange[0] = nl
			newRange[1] = nr
			ret = append(ret, newRange)
		}
	}
	return ret
}
