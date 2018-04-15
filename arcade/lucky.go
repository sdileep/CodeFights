package arcade

// IsLucky returns whether a ticket is lucky
func IsLucky(n int) bool {

	var res []int
	// get the digits
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}

	// loop through array and track every second element differently
	var first, second int
	for i := 0; i < len(res)/2; i++ {
		first += res[i]
	}
	for i := len(res) / 2; i < len(res); i++ {
		second += res[i]
	}

	return first == second
}
