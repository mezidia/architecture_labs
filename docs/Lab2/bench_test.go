package lab2

import (
	"fmt"
	"testing"
)

var (
	startStatement = "+ - / * 1 2 3 4 5"
)

func BenchmarkPrefix(b *testing.B) {
	iterationsNum := 0

	for k := 0; k < 20; k++ {
		input := startStatement
		iterationsNum += 2000

		for i := 0; i < iterationsNum; i++ {
			input = "+ " + input
			input = input + " "
			input = input + startStatement
		}

		b.Run(fmt.Sprintf("len=%d", iterationsNum), func(b *testing.B) {
			cntRes, err := Prefix(input)

			if err != nil {
				fmt.Println(err)
				fmt.Println("Result is ", cntRes)
			}
		})
	}
}
