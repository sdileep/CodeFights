package arcade

// AllLongestStrings returns the longest substrings in the list
func AllLongestStrings(inputArray []string) []string {
	var res []string
	max := len(inputArray[0])
	for _, e := range inputArray {
		if len(e) > max {
			max = len(e)
		}
	}
	for _, e := range inputArray {
		if len(e) == max {
			res = append(res, e)
		}
	}
	return res
}
