package lab2

import (
	"fmt"
	"testing"
)

var cntRes int
var input string
var err error

func BenchmarkPrefix(b *testing.B) {
	l := 0
	expr := "+ - / * 1 2 3 4 5"
	for k := 0; k < 20; k++ {
		input = expr
		l += 2000
		for i := 0; i < l; i++ {
			input = "+ " + input
			input = input + " "
			input = input + expr
		}

		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cntRes, err = Prefix(input)
		})
	}
}
