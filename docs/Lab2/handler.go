package lab2

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	// TODO: Add necessary fields.
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	inputExpression := flag.String("e", "", "Expression to compute")
	inputFile := flag.String("f", "", "File with expression")
	outputFile := flag.String("o", "", "File with result")

	flag.Parse()

	expressionToUse := ""
	if len(*inputExpression) > 0 {
		expressionToUse = *inputExpression

	} else if len(*inputFile) > 0 {
		input, err := os.ReadFile(*inputFile)

		if err != nil {
			log.Fatal(err)
		}

		expressionToUse = string(input)
	}

	res, err := Prefix(expressionToUse)
	if err == nil {
		if len(*outputFile) > 0 {
			resInString := strconv.Itoa(res)
			data := []byte(resInString)

			os.WriteFile(*outputFile, data, 0666)
		} else {
			fmt.Fprintln(ch.Output, res)
		}
	} else {
		fmt.Fprintln(ch.Output, err)
	}
	return nil
}
