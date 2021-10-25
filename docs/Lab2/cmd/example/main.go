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

	var in *strings.Reader
	if len(*inputExpression) > 0 {
		in = strings.NewReader(*inputExpression)
	}
	else

	handler := &lab2.ComputeHandler{
		Input:  in,
		Output: os.Stdout,
	}
	err := handler.Compute()
	if err != nil {
		fmt.Println(err)
	}
}
