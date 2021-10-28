package lab2

import (
	"bytes"
	"io"
	"strconv"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(ch.Input)
	expression := buf.String()

	res, err := Prefix(expression)
	if err == nil {
		resInString := strconv.Itoa(res)
		data := []byte(resInString)
		ch.Output.Write(data)
	}
	return nil
}
