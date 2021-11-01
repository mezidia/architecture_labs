package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	input := "+ 3 2"
	inputReader := strings.NewReader(input)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  inputReader,
		Output: output,
	}
	err := handler.Compute()
	assert.Nil(t, err)
}
