package arcade_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testLuckyCases = []struct {
	in       int
	expected bool
}{
	{1230, true},
	{239017, false},
}

func TestIsLucky(t *testing.T) {
	for _, test := range testLuckyCases {
		got := arcade.IsLucky(test.in)
		if got != test.expected {
			t.Fatalf("Got %v, expected %v. Input %d", got, test.expected, test.in)
		}
	}
}
