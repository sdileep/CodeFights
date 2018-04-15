package arcade_test

import (
	"reflect"
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testLongestSubstringCases = []struct {
	in       []string
	expected []string
}{
	{[]string{"aba", "aa", "ad", "vcd", "aba"}, []string{"aba", "vcd", "aba"}},
	{[]string{"abc", "eeee", "abcd", "dcd"}, []string{"eeee", "abcd"}},
}

func TestAllLongestSubstring(t *testing.T) {
	for _, test := range testLongestSubstringCases {
		got := arcade.AllLongestStrings(test.in)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Got %v, expected %v", got, test.expected)
		}
	}
}
