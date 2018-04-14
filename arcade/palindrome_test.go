package arcade_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testPalindromeCases = []struct {
	in       string
	expected bool
}{
	{"aabaa", true},
	{"abac", false},
	{"a", true},
}

func TestCheckPalindrome(t *testing.T) {
	for _, test := range testPalindromeCases {
		got := arcade.CheckPalindrome(test.in)
		if got != test.expected {
			t.Fatalf("Got %v, expected %v", got, test.expected)
		}
	}
}
