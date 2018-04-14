package arcade_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testShapeCases = []struct {
	in       int
	expected int
}{
	{1, 1},
	{2, 5},
	{3, 13},
	{4, 25},
}

func TestShapeArea(t *testing.T) {
	for _, test := range testShapeCases {
		got := arcade.ShapeArea(test.in)
		if got != arcade.ShapeArea(test.in) {
			t.Fatalf("Got %d, expected %d\n", got, test.expected)
		}
	}
}
