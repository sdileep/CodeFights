package core_test

import (
	"testing"

	"bitbucket.org/ffxblue/CodeFights/core"
)

var testAddCases = []struct {
	in       int
	expected int
}{
	{29, 11},
	{48, 12},
}

func TestAddTwoDigits(t *testing.T) {
	for _, test := range testAddCases {
		got := core.AddTwoDigits(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d from %d", got, test.expected, test.in)
		}
	}
}
