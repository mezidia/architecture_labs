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
	com := lab2.ComputeHandler{
		In:  strings.NewReader("Main reader"),
		Out: 
	}
	com.Compute()
}
