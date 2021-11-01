package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/mezidia/architecture_labs/tree/main/docs/Lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression")
	outputFile      = flag.String("o", "", "File with result")
)

func main() {
	flag.Parse()

	if (*inputExpression != "" && *inputFile != "") ||
		(*inputExpression == "" && *inputFile == "") {
		fmt.Println("unexpected input")
	} else {
		var input io.Reader
		var output io.Writer

		if len(*inputExpression) > 0 {
			input = strings.NewReader(*inputExpression)
		} else {
			inputExpressionFromFile, err := os.ReadFile(*inputFile)
			if err != nil {
				fmt.Println(err)
			}
			input = bytes.NewReader(inputExpressionFromFile)
		}

		if len(*outputFile) > 0 {
			outputToFile, err := os.Create(*outputFile)
			if err != nil {
				fmt.Println(err)
			}
			output = outputToFile
		} else {
			output = os.Stdout
		}

		handler := &lab2.ComputeHandler{
			Input:  input,
			Output: output,
		}
		err := handler.Compute()
		if err != nil {
			fmt.Println(err)
		}
	}
}
