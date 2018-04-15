package arcade

import "sort"

// SortByHeight returns a sorted array with trees in between
func SortByHeight(a []int) []int {
	var wot []int
	for _, e := range a {
		if e != -1 {
			wot = append(wot, e)
		}
	}
	sort.Ints(wot)
	var wc int
	for i, e := range a {
		if e != -1 {
			a[i] = wot[wc]
			wc++
		}
	}

	return a
}
