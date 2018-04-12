package basic_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/basic"
)

var testProductCases = []struct {
	in       []int
	mod      int
	expected int
}{
	{
		[]int{1, 2, 3, 4},
		12,
		2,
	},
	{
		[]int{2, 100},
		24,
		6,
	},
}

func TestProductExceptSelf(t *testing.T) {
	for _, test := range testProductCases {
		got := basic.ProductExceptSelf(test.in, test.mod)
		if got != test.expected {
			t.Fatalf("Got %d, want %d", got, test.expected)
		}
	}
}
