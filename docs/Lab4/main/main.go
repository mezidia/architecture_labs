package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mezidia/architecture_labs/tree/main/docs/Lab4/engine"
)

func main() {

	loop := new(engine.Loop)

	loop.Start()
	inputFile := "examples.txt"
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			loop.Post(cmd)
		}
	}

	loop.AwaitFinish()
	loop.Post(&engine.PrintCommand{Arg: "Should end!"})
}

func parse(command string) engine.Command {
	parts := strings.Fields(command)
	if parts[0] == "print" {
		return (&engine.PrintCommand{Arg: parts[1]})
	} else if parts[0] == "add" {
		firstNum, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
		}
		secondNum, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println(err)
		}
		return (&engine.AddCommand{A: firstNum, B: secondNum})
	} else {
		return (&engine.PrintCommand{Arg: "Something wrong Happend"})
	}
}
