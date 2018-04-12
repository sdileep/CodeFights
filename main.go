package main

import (
	"fmt"

	"github.com/sunilgopinath/codefights/basic"
)

func main() {
	h := basic.ProductExceptSelf([]int{1, 2, 3, 4}, 12)
	fmt.Println(h)
}
