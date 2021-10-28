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
		var inputFromCMD *strings.Reader
		var inputFromFile io.Reader
		var input io.Reader
		var output io.Writer

		if len(*inputExpression) > 0 {
			inputFromCMD = strings.NewReader(*inputExpression)
			input = inputFromCMD
		} else {
			inputExpressionFromFile, err := os.ReadFile(*inputFile)
			if err != nil {
				fmt.Println(err)
			}
			// convert byte slice to io.Reader
			inputFromFile = bytes.NewReader(inputExpressionFromFile)
			input = inputFromFile
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
