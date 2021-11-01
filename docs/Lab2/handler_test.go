package lab2

import (
	"testing"

	lab2 "github.com/mezidia/architecture_labs/tree/main/docs/Lab2"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	input := "+ 3 2"
	output := ""
	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	assert.NotNil(err)
}
