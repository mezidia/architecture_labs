package lab2

import (
	"fmt"
	"testing"
)

var startStatement string = "+ - / * 1 2 3 4 5"
var cntRes int
var err error

func BenchmarkPrefix(b *testing.B) {
	const baseLength = 2500

	for i := 0; i < 25; i++ {
		input := startStatement
		iterationsNum := baseLength * (i + 1)

		for j := 0; j < iterationsNum; j++ {
			input = "+ " + input
			input = input + " "
			input = input + startStatement
		}

		b.Run(fmt.Sprintf("len=%d", iterationsNum), func(b *testing.B) {
			cntRes, err = Prefix(input)
		})
	}
}
