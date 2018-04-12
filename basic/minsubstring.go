package basic

import (
	"math"
)

// MinSubstringWithAllChars returns the min substring of s containing all t
func MinSubstringWithAllChars(s string, t string) string {

	// shortcuts
	if s == t {
		return s
	}
	if len(t) == 1 {
		return t // guaranteed to exist according to spec
	}

	// Algorithm

	// turn t into an fast lookup structure
	m := mapify(t)

	// iterate over s
	rr := []rune(s)

	// initial value to hold min min substring
	msc := math.MaxInt32
	ms := ""

	// var start int
	var ss string
	for i := 0; i < len(rr)-len(t); i++ {
		r := rr[i]
		if _, ok := m[r]; ok {
			ss = search(i, rr, mapify(t)) // need to pass in a reference to the map
			if ss != "" {
				if len(ss) < msc {
					ms = ss
					msc = len(ss)
				}
			}
		}
	}
	return ms
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func search(start int, rr []rune, m map[rune]int) string {
	var res []rune
	var fc int
	for i := start; i < len(rr); i++ {
		r := rr[i]
		res = append(res, r)
		if _, ok := m[r]; ok {
			// skip over already found values
			if m[r] == 0 {
				fc++
			}
			m[r]++
		}
		if fc == len(m) {
			return string(res)
		}
	}
	return "" // can get here if we don't find everything
}

func mapify(t string) map[rune]int {
	m := make(map[rune]int)
	rr := []rune(t)
	for _, r := range rr {
		m[r] = 0
	}
	return m
}
