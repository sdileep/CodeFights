package backtracking

// SumSubsets returns the subsets which add to a sum
func SumSubsets(arr []int, num int) [][]int {
	var curr []int
	var processed [][]int
	sumSubsetsHelper(arr, curr, &processed, num, 0)
	return processed
}

func sumSubsetsHelper(arr, curr []int, processed *[][]int, num, lastSeen int) {
	if num == 0 {
		*processed = append(*processed, append([]int{}, curr...))
		return
	}
	for i := 0; i < len(arr); i++ {
		// choose element
		chosen := arr[i]
		if chosen == lastSeen {
			continue
		}
		curr = append(curr, chosen)
		if num-chosen >= 0 {
			sumSubsetsHelper(arr[i+1:], curr, processed, num-chosen, lastSeen)
		}
		lastSeen = arr[i]
		curr = curr[:len(curr)-1]
	}
}
