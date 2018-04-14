package dp_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/dp"
)

var testClimbingCases = []struct {
	in       int
	expected int
}{
	{
		1,
		1,
	},
	{
		2,
		2,
	},
	{
		26,
		196418,
	},
}

func TestClimbingStairsNaive(t *testing.T) {
	for _, test := range testClimbingCases {
		got := dp.ClimbingStairsNaive(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d", got, test.expected)
		}
	}
}

func TestClimbingStairsMemo(t *testing.T) {
	for _, test := range testClimbingCases {
		got := dp.ClimbingStairsMemo(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d", got, test.expected)
		}
	}
}

func TestClimbingStairsDP(t *testing.T) {
	for _, test := range testClimbingCases {
		got := dp.ClimbingStairsDP(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d", got, test.expected)
		}
	}
}
