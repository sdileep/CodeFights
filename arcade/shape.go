package arcade

//ShapeArea returns the n
func ShapeArea(n int) int {
	// formula is 4 * i + shapeArea(n-1)
	if n == 1 {
		return n
	}
	curr := 0
	prev := 1
	for i := 1; i < n; i++ {
		curr = 4*i + prev
		prev = curr
	}
	return curr
}
