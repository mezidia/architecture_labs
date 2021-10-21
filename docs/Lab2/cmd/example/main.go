package main

import (
	"flag"
	"fmt"

	lab2 "github.com/mezidia/architecture_labs/tree/main/docs/Lab2"
)

// var (
// 	inputExpression = flag.String("e", "", "Expression to compute")
// 	// TODO: Add other flags support for input and output configuration.
// )

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, _ := lab2.Prefix("+ 11 20")
	fmt.Println(res)
}