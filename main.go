package main

import (
	"fmt"

	"github.com/sunilgopinath/codefights/arcade"
)

func main() {
	// h := basic.ProductExceptSelf([]int{1, 2, 3, 4}, 12)
	// fmt.Println(h)
	h := arcade.MakeArrayConsecutive2([]int{6, 2, 3, 8})
	fmt.Println(h)
}
