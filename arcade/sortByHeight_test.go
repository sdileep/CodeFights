package arcade_test

import (
	"reflect"
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testSortCases = []struct {
	in       []int
	expected []int
}{
	{
		[]int{-1, 150, 190, 170, -1, -1, 160, 180},
		[]int{-1, 150, 160, 170, -1, -1, 180, 190},
	},
}

func TestSortByHeight(t *testing.T) {
	for _, test := range testSortCases {
		got := arcade.SortByHeight(test.in)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Got %v, expected %v from %v", got, test.expected, test.in)
		}
	}
}
