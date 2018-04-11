package main

import (
	"fmt"

	"github.com/sunilgopinath/codefights/basic"
)

func main() {
	h := basic.SumInRange([]int{3, 0, -2, 6, -3, 2}, [][]int{{0, 2}, {2, 5}, {0, 5}})
	fmt.Println(h)
}
