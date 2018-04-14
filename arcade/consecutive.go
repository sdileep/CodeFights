package arcade

import (
	"sort"
)

//MakeArrayConsecutive2 returns the number of missing numbers
func MakeArrayConsecutive2(arr []int) int {
	sort.Ints(arr)
	var c int
	for i := 1; i < len(arr); i++ {
		c += arr[i] - arr[i-1] - 1
	}
	return c
}
