package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

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
	if len(*inputExpression) > 0 {
		expressionToUse = *inputExpression
	} else if len(*inputFile) > 0 {
		input, err := os.ReadFile(*inputFile)

		if err != nil {
			log.Fatal(err)
		}

		expressionToUse = string(input)
	}

	res, err := lab2.Prefix(expressionToUse)
	if err == nil {
		if len(*outputFile) > 0 {
			data := strconv.Itoa(res)

			bytes := []byte(data)
			os.WriteFile(*outputFile, bytes, 0666)
		} else {
			fmt.Println(res)
		}
	}
}
