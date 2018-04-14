package arcade

//CheckPalindrome returns whether a string is a palindrome
func CheckPalindrome(inputString string) bool {
	i := 0
	j := len(inputString) - 1
	rr := []rune(inputString)
	for i <= j {
		if rr[i] != rr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
