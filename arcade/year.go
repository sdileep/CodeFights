package arcade

// CenturyFromYear returns the century given the year as input
func CenturyFromYear(year int) int {
	c := year / 100
	r := year % 100
	if c < 1 {
		return 1
	} else if r > 0 {
		return c + 1
	}
	return c
}
