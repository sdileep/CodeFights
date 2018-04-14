package arcade

import "math"

// AdjacentElementsProduct returns the max product of an adjacent pair
func AdjacentElementsProduct(arr []int) int {
	max := math.MinInt32
	for i := 1; i < len(arr); i++ {
		p := arr[i] * arr[i-1]
		if p > max {

			max = p
		}
	}
	return max
}
