package main

import (
	"flag"
	"fmt"
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
		fmt.Errorf("unexpected input")
	} else {
		var inputFromCMD *strings.Reader
		var inputFromFile os.File

		if len(*inputExpression) > 0 {
			inputFromCMD = strings.NewReader(*inputExpression)
		}

		handler := &lab2.ComputeHandler{
			Input:  inputFromCMD,
			Output: os.Stdout,
		}
		err := handler.Compute()
		if err != nil {
			fmt.Println(err)
		}
	}
}
