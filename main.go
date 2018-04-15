package main

import (
	"fmt"

	"github.com/sunilgopinath/codefights/arcade"
)

func main() {
	// h := basic.ProductExceptSelf([]int{1, 2, 3, 4}, 12)
	// fmt.Println(h)
	h := arcade.MatrixElementsSum([][]int{{1, 1, 1, 0}, {0, 5, 0, 1}, {2, 1, 3, 10}})
	fmt.Println(h)
}
