package arcade_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testMatrixCases = []struct {
	in       [][]int
	expected int
}{
	{
		[][]int{{1, 1, 1, 0}, {0, 5, 0, 1}, {2, 1, 3, 10}},
		9,
	},
}

func TestMatrixElementsSum(t *testing.T) {
	for _, test := range testMatrixCases {
		got := arcade.MatrixElementsSum(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d", got, test.expected)
		}
	}
}
