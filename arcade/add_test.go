package arcade_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testAddCases = []struct {
	in1      int
	in2      int
	expected int
}{
	{1, 6, 7},
	{2, 3, 5},
	{1, 2, 3},
}

func TestAdd(t *testing.T) {
	for _, test := range testAddCases {
		got := arcade.Add(test.in1, test.in2)
		if got != test.expected {
			t.Fatalf("Got %d, want %d\n", got, test.expected)
		}
	}
}
