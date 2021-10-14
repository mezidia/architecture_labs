package lab2

import (
	"math"
	"unicode"
)

func isDigit(symbol rune) bool {
	return unicode.IsDigit(symbol)
}

// TODO: document this function.
// Prefix converts
func Prefix(input string) (int, error) {
	var stack []int

	// "+ 2 2"

	for j := len(input) - 1; j >= 0; j-- {
		if rune(input[j]) == ' ' {
			continue
		}
		if isDigit(rune(input[j])) {
			num := 0
			i := j
			for j < len(input) && isDigit(rune(input[j])) {
				j--
			}
			j++

			for k := j; k <= i; k++ {
				num = num*10 + int(input[k]-'0')
			}
			stack = append(stack, num)
		} else {
			// length := len(stack)
			o1 := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]
			o2 := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]
			// fmt.Println(o1)
			// fmt.Println(o2)
			// fmt.Println(stack)
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
