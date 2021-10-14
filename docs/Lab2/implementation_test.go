package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix_Easy_First(t *testing.T) {
	expected, err := Prefix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		var actual = 11
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_Addition(t *testing.T) {
	expected, err := Prefix("+ 2 2")
	if assert.Nil(t, err) {
		var actual = 4
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_1Addition(t *testing.T) {
	expected, err := Prefix("+ 11 2")
	if assert.Nil(t, err) {
		var actual = 13
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_Subtraction(t *testing.T) {
	expected, err := Prefix("- 2 2")
	if assert.Nil(t, err) {
		var actual = 0
		assert.Equal(t, actual, expected)
	}
}

// func TestPrefix_Division(t *testing.T) {
// 	expected, err := Prefix("/ 16 2")
// 	if assert.Nil(t, err) {
// 		var actual = 8
// 		assert.Equal(t, actual, expected)
// 	}
// }

func TestPrefix_Multiplication(t *testing.T) {
	expected, err := Prefix("* 3 4")
	if assert.Nil(t, err) {
		var actual = 12
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_Power(t *testing.T) {
	expected, err := Prefix("^ 2 5")
	if assert.Nil(t, err) {
		var actual = 32
		assert.Equal(t, actual, expected)
	}
}

func ExamplePrefix() {
	expected, _ := Prefix("+ 2 2")
	fmt.Println(expected)

	// Output:
	// 4
}
