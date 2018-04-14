package dp

// ClimbingStairsNaive calculates distinct ways in climbing a staircase
func ClimbingStairsNaive(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	if n == 2 {
		return 2
	}

	// You can reach the current step either by taking 1 step OR 2 steps
	return ClimbingStairsNaive(n-1) + ClimbingStairsNaive(n-2)
}

// ClimbingStairsMemo calculates distinct ways in climbing a staircase
func ClimbingStairsMemo(n int) int {
	m := make(map[int]int)
	return climbingStairsHelper(n, m)
}

func climbingStairsHelper(n int, m map[int]int) int {
	if _, ok := m[n]; ok {
		return m[n]
	}
	if n == 0 || n == 1 {
		return n
	}

	if n == 2 {
		return 2
	}
	result := climbingStairsHelper(n-1, m) + climbingStairsHelper(n-2, m)

	m[n] = result
	return result
}

// ClimbingStairsDP calculates distinct ways in climbing a staircase
func ClimbingStairsDP(n int) int {
	return climbingStairsDPHelper(n + 1)
}

func climbingStairsDPHelper(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	prevPrev := 0
	prev := 1
	current := 0
	for i := 1; i < n; i++ {
		current = prev + prevPrev
		prevPrev = prev
		prev = current
	}
	return current
}
