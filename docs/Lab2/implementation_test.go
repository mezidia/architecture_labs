package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		var expected = 32
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

func TestPrefix_SubtractionWithDozens(t *testing.T) {
	actual, err := Prefix("- 11 20")
	if assert.Nil(t, err) {
		var expected = -9
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_SubtractionWithHundreds(t *testing.T) {
	actual, err := Prefix("- 343 590")
	if assert.Nil(t, err) {
		var expected = -247
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_SubtractionWithThousands(t *testing.T) {
	actual, err := Prefix("- 3449 2910")
	if assert.Nil(t, err) {
		var expected = 539
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_Division(t *testing.T) {
	actual, err := Prefix("/ 4 2")
	if assert.Nil(t, err) {
		var expected = 2
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_DivisionWithDozens(t *testing.T) {
	actual, err := Prefix("/ 30 10")
	if assert.Nil(t, err) {
		var expected = 3
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_DivisionWithHundreds(t *testing.T) {
	actual, err := Prefix("/ 600 150")
	if assert.Nil(t, err) {
		var expected = 4
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_DivisionWithThousands(t *testing.T) {
	actual, err := Prefix("/ 9999 1111")
	if assert.Nil(t, err) {
		var expected = 9
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

func TestPrefix_MultiplicationWithDozens(t *testing.T) {
	actual, err := Prefix("* 30 10")
	if assert.Nil(t, err) {
		var expected = 300
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_MultiplicationWithHundreds(t *testing.T) {
	actual, err := Prefix("* 600 150")
	if assert.Nil(t, err) {
		var expected = 90000
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_MultiplicationWithThousands(t *testing.T) {
	actual, err := Prefix("* 9999 1111")
	if assert.Nil(t, err) {
		var expected = 11108889
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

func TestPrefix_FirstEasy(t *testing.T) {
	actual, err := Prefix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		var expected = 11
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_SecondEasy(t *testing.T) {
	actual, err := Prefix("* - 5 6 7")
	if assert.Nil(t, err) {
		var expected = -7
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_ThirdEasy(t *testing.T) {
	actual, err := Prefix("* + 1 2 + 3 4")
	if assert.Nil(t, err) {
		var expected = 21
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_FirstHard(t *testing.T) {
	actual, err := Prefix("- * / 15 - 7 + 1 1 3 + 2 + 1 1")
	if assert.Nil(t, err) {
		var expected = 5
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_SecondHard(t *testing.T) {
	actual, err := Prefix("+ + - - + 4 * 3 5 ^ 10 3 1 * / 120 5 3 2")
	if assert.Nil(t, err) {
		var expected = -908
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_ThirdHard(t *testing.T) {
	actual, err := Prefix("- + - + - * / 4 2 6 3 221 ^ 321 2 / 3 3 * 98 2")
	if assert.Nil(t, err) {
		var expected = -103006
		assert.Equal(t, actual, expected)
	}
}

func TestPrefix_OnlyOneDivisionError(t *testing.T) {
	_, err := Prefix("/")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("unexpected input"), err)
	}
}

func TestPrefix_OnlyOnePowError(t *testing.T) {
	_, err := Prefix("^")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("unexpected input"), err)
	}
}

func TestPrefix_OnlyTwoOperandsError(t *testing.T) {
	_, err := Prefix("- +")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("unexpected input"), err)
	}
}

func TestPrefix_BlankError(t *testing.T) {
	_, err := Prefix("")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("unexpected input"), err)
	}
}

func TestPrefix_EmptyError(t *testing.T) {
	_, err := Prefix(" ")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("blank input"), err)
	}
}

func TestPrefix_OneLetterError(t *testing.T) {
	_, err := Prefix("f")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("unexpected symbol"), err)
	}
}

func TestPrefix_LettersError(t *testing.T) {
	_, err := Prefix("DFSDFSDF")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("unexpected symbol"), err)
	}
}

func TestPrefix_MoreDigitsThanOperandsError(t *testing.T) {
	_, err := Prefix("+ 3 3 3")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("unexpected input"), err)
	}
}

func TestPrefix_BigDigitsMoreThanOperandsError(t *testing.T) {
	_, err := Prefix("+ 33 322 3111")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("unexpected input"), err)
	}
}

func ExamplePrefix() {
	actual, _ := Prefix("+ 2 2")
	fmt.Println(actual)

	// Output:
	// 4
}
