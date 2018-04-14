package main

import (
	"fmt"

	"github.com/sunilgopinath/codefights/arcade"
)

func main() {
	// h := basic.ProductExceptSelf([]int{1, 2, 3, 4}, 12)
	// fmt.Println(h)
	h := arcade.AdjacentElementsProduct([]int{-23, 4, -3, 8, -12})
	fmt.Println(h)
}
