package arcade

// MatrixElementsSum returns the total cost of all rooms
func MatrixElementsSum(matrix [][]int) int {
	m := make(map[int]int)
	var total int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				m[j] = 1
			}
			if _, ok := m[j]; ok {
				continue
			}
			total += matrix[i][j]
		}
	}
	return total
}
