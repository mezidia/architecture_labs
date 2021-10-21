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
	if input == " " {
		panic("blank input")
	}

	var numbers, chars int
	// 48 - 57 = numbers
	// ^ - 94; * - 42; + - 43; - - 45; / - 47
	for index := 0; index < len(input); index++ {
		symbol := input[index]
		if 48 <= symbol && symbol <= 57 {
			// found a number
			numbers++
		} else if symbol == 94 || symbol == 42 || symbol == 43 || symbol == 45 || symbol == 47 {
			// found an ariphmetic operator
			chars++
		} else if symbol != 32 { // not space
			panic("unexpected symbol")
		}
	}
	if numbers-1 != chars {
		panic("unexpected input")
	}

	var stack []int
	var asciiOffset byte = 48
	var multiplyOffset int = 10

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
				num = multiplyOffset*num + int(input[k]-asciiOffset) //was -'0'
			}
			stack = append(stack, num)

		} else {
			firstOperand := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]

			secondOperand := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]

			switch rune(input[j]) {
			case '+':
				result := firstOperand + secondOperand
				stack = append(stack, result)
				break
			case '-':
				result := firstOperand - secondOperand
				stack = append(stack, result)
				break
			case '*':
				result := firstOperand * secondOperand
				stack = append(stack, result)
				break
			case '/':
				result := firstOperand / secondOperand
				stack = append(stack, result)
				break
			case '^':
				result := int(math.Pow(float64(firstOperand), float64(secondOperand)))
				stack = append(stack, result)
				break
			}
		}
	}
	return stack[0], nil
}
