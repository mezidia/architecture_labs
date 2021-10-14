package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix(t *testing.T) {
	expected, err := Prefix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		var actual = 11
		assert.Equal(t, actual, expected)
	}
}

func ExamplePrefix() {
	expected, _ := Prefix("+ 2 2")
	fmt.Println(expected)

	// Output:
	// 4
}
