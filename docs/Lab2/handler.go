package lab2

import (
	"bytes"
	"fmt"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(ch.Input)

	res, err := Prefix(buf.String())
	if err == nil {
		// 	if len(*outputFile) > 0 {
		// 		resInString := strconv.Itoa(res)
		// 		data := []byte(resInString)

		// 		os.WriteFile(*outputFile, data, 0666)
		// 	} else {
		// 		fmt.Fprintln(ch.Output, res)
		// 	}
		// } else {
		// 	fmt.Fprintln(ch.Output, err)
	}
	fmt.Println(res)
	return nil
}
