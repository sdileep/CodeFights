package basic_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/basic"
)

var testMinSubstringCases = []struct {
	s        string
	t        string
	expected string
}{
	{
		"adobecodebanc",
		"abc",
		"banc",
	},
	{
		"",
		"",
		"",
	},
	{
		"tvdsxcqsnoeccaurocnk",
		"acqt",
		"tvdsxcqsnoecca",
	},
	{
		"xgajymplpvftjwjqomhbnutorgysaf",
		"j",
		"j",
	},
}

func TestMinSubstringWithAllChars(t *testing.T) {
	for _, test := range testMinSubstringCases {
		got := basic.MinSubstringWithAllChars(test.s, test.t)
		if got != test.expected {
			t.Fatalf("Got %s, want %s", got, test.expected)
		}
	}
}
