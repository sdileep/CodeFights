package core

// AddTwoDigits returns the sum of the digits
func AddTwoDigits(n int) int {
	var total int
	for n > 0 {
		total += n % 10
		n /= 10
	}
	return total
}
