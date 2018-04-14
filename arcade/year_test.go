package arcade_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testYearCases = []struct {
	y        int
	expected int
}{
	{8, 1},
	{1905, 20},
	{1700, 17},
}

func TestCenturyFromYear(t *testing.T) {
	for _, test := range testYearCases {
		got := arcade.CenturyFromYear(test.y)
		if got != test.expected {
			t.Fatalf("Got %d, want %d", got, test.expected)
		}
	}
}
