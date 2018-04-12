package basic

// ProductExceptSelf calculates products without self
func ProductExceptSelf(nums []int, m int) int {

	/*
	  ex [1,2,3,4,5]
	  the resulting slice would be [2*3*4*5, 1*3*4*5, 1*2*3*5, 1*2*3*4]
	*/

	// result slice
	res := make([]int, len(nums))

	// here we are going to calculate the product before index of each element
	pf := 1
	for i, n := range nums {
		res[i] = pf
		pf *= n
		pf %= m // modulo as soon as possible so we do not get very large numbers
	}

	// now we are going to multiply each element by the number
	// we would get by calculating after the element
	pf = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= pf
		pf *= nums[i]
		pf %= m
	}

	// as per question, get sum
	var t int
	for _, e := range res {
		t += e
	}

	// return sum %
	return t % m
}
