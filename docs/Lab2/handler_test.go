package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Compute(t *testing.T) {
	input := "+ 3 2"
	var expected = "5"
	inputReader := strings.NewReader(input)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  inputReader,
		Output: output,
	}
	err := handler.Compute()
	assert.Equal(t, output.String(), expected)
	assert.Nil(t, err)
}

func Test_Compute_Error(t *testing.T) {
	input := "+ 3 2 2"
	inputReader := strings.NewReader(input)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  inputReader,
		Output: output,
	}
	err := handler.Compute()
	assert.Equal(t, output.String(), "")
	assert.NotNil(t, err)
	// You are durik
}
