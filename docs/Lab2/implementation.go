package lab2

import (
	"math"
	"unicode"
)

func isNumber(symbol rune) bool {
	return unicode.IsDigit(symbol)
}

// TODO: document this function.
// Prefix converts
func Prefix(input string) (int, error) {
	var stack []int

	for j := len(input) - 1; j >= 0; j-- {
		if rune(input[j]) == ' ' {
			continue
		}

		if isNumber(rune(input[j])) {
			num := 0
			i := j
			for j < len(input) && isNumber(rune(input[j])) {
				j--
			}
			j++

			for k := j; k <= i; k++ {
				num = num*10 + int(input[k]-'0')
			}
			stack = append(stack, num)
		} else {
			firstOperand := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]

			secondOperand := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]

			switch rune(input[j]) {
			case '+':
				stack = append(stack, o1+o2)
				break
			case '-':
				stack = append(stack, o1-o2)
				break
			case '*':
				stack = append(stack, o1*o2)
				break
			case '/':
				stack = append(stack, o1/o2)
				break
			case '^':
				stack = append(stack, int(math.Pow(float64(o1), float64(o2))))
				break
			}
		}
	}
	return stack[0], nil
}
