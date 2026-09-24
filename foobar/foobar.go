package foobar

import (
	"fmt"
)

// Sequence returns a slice of strings containing numbers from 1 to length
// where multiples of 3 are replaced by "foo", multiples of 5 by "bar",
// and multiples of both 3 and 5 by "foobar".
func Sequence(length int) ([]string, error) {
	if length < 0 {
		return nil, fmt.Errorf("length cannot be negative")
	}

	seq := make([]string, length)
	for i := range seq {
		n := i + 1
		switch {
		case n%3 == 0 && n%5 == 0:
			seq[i] = "foobar"
		case n%5 == 0:
			seq[i] = "bar"
		case n%3 == 0:
			seq[i] = "foo"
		default:
			seq[i] = fmt.Sprintf("%d", n)
		}
	}
	return seq, nil
}