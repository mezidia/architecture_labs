package lab4

import (
	"fmt"
	"testing"

	"github.com/mezidia/architecture_labs/tree/main/docs/Lab4/engine"
)

var txt = "print script"
var cntRes engine.Command

func BenchmarkCount(b *testing.B) {
	const baseLen = 7500
	for i := 0; i < 16; i++ {
		input := txt
		l := baseLen * (i + 1)

		for j := 0; j < l; j++ {
			input = input + "sadstorylinesadstorylinesadstorylinesadstoryline"
		}

		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cntRes = Parse(input)
		})
	}
}
