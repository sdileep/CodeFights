package arcade_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testConsecutiveCases = []struct {
	in       []int
	expected int
}{
	{[]int{6, 2, 3, 8}, 3},
	{[]int{0, 3}, 2},
}

func TestMakeArrayConsecutive2(t *testing.T) {
	for _, test := range testConsecutiveCases {
		got := arcade.MakeArrayConsecutive2(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d", got, test.expected)
		}
	}
}
