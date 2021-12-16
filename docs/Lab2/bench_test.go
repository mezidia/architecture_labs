package lab2

import (
	"fmt"
	"math/rand"
	"testing"
)

var cntRes int

func BenchmarkCount(b *testing.B) {
	const baseLen = 5
	for i := 0; i < 20; i++ {
		l := baseLen * 1 << i
		in := make([]int, l)
		for k := 0; k < l; k++ {
			in[k] = rand.Intn(10)
		}
		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cntRes = Count(make([]int, l), 0)
		})
	}
}
