package main

import (
	"debug/elf"
	"flag"
	"fmt"
	"io"

	lab2 "github.com/mezidia/architecture_labs/tree/main/docs/Lab2"
)

// var (
// 	inputExpression = flag.String("e", "", "Expression to compute")
// 	// TODO: Add other flags support for input and output configuration.
// )

func main() {
	// var (
	// 	inputExpression = flag.String("e", "", "Expression to compute")
	// 	// TODO: Add other flags support for input and output configuration.
	// )

	inputExpression := flag.String("e", "", "Expression to compute")
	inputFile := flag.String("f", "", "File with expression")
	outputFile := flag.String("o", "", "File with result")

	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	expressionToUse := ""

	if inputExpression != nil {
		expressionToUse = *inputExpression
	}
	else if inputFile != nil {
		expressionToUse = *inputExpression
	}

	res, err := lab2.Prefix(expressionToUse)
	if err == nil {
		fmt.Println(res)
	}
}
