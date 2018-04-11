package basic

// SumInRange looks for sums of queries in ranges
func SumInRange(nums []int, queries [][]int) int {
	var total int

	// cumulative sum
	cs := cumulativeSum(nums)

	// now iterate over queries slice
	for _, e := range queries {
		if e[0] == 0 {
			// we start from the beginning so can just return the
			// array value of the highest range ie e[1]
			total += cs[e[1]]
		} else {
			// we have to eliminate from the top and bottom
			total += cs[e[1]] - cs[e[0]-1]
		}
	}
	total %= 1000000007
	if total < 0 {
		total += 1000000007
	}
	return total % 1000000007
}

func cumulativeSum(nums []int) []int {
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] + nums[i]
	}
	return arr
}
