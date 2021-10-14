package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix_Easy_First(t *testing.T) {
	actual, err := Prefix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		var expected = 11
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_Addition(t *testing.T) {
	actual, err := Prefix("+ 2 2")
	if assert.Nil(t, err) {
		var expected = 4
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_AdditionWithDozens(t *testing.T) {
	actual, err := Prefix("+ 11 20")
	if assert.Nil(t, err) {
		var expected = 31
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_AdditionWithHundreds(t *testing.T) {
	actual, err := Prefix("+ 343 590")
	if assert.Nil(t, err) {
		var expected = 933
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_AdditionWithThousands(t *testing.T) {
	actual, err := Prefix("+ 3449 2910")
	if assert.Nil(t, err) {
		var expected = 6359
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_Subtraction(t *testing.T) {
	actual, err := Prefix("- 2 2")
	if assert.Nil(t, err) {
		var expected = 0
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_Division(t *testing.T) {
	actual, err := Prefix("/ 16 2")
	if assert.Nil(t, err) {
		var expected = 8
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_Multiplication(t *testing.T) {
	actual, err := Prefix("* 3 4")
	if assert.Nil(t, err) {
		var expected = 12
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_Power(t *testing.T) {
	actual, err := Prefix("^ 2 5")
	if assert.Nil(t, err) {
		var expected = 32
		assert.Equal(t, actual, expected)
	}
}

func ExamplePrefix() {
	actual, _ := Prefix("+ 2 2")
	fmt.Println(actual)

	// Output:
	// 4
}
