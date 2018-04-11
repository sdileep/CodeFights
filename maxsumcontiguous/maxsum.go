package maxsumcontiguous

// ArrayMaxConsecutiveSum2 returns the longest continuous subarray sum
func ArrayMaxConsecutiveSum2(aa []int) int {
	lmax := aa[0]
	omax := aa[0]
	for i := 1; i < len(aa); i++ {
		lmax = max(aa[i], lmax+aa[i])
		omax = max(omax, lmax)
	}
	return omax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
