package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/mezidia/architecture_labs/tree/main/docs/Lab4/engine"
)

var (
	loopEnd      = "Should end!"
	inputFile    = "examples.txt"
	syntaxError  = "SYNTAX ERROR: "
	commandError = "unknown command"
	printCommand = "print"
	addCommand   = "add"
)

func main() {

	loop := new(engine.Loop)

	loop.Start()
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
	loop.Post(&engine.PrintCommand{Arg: loopEnd})
}

func parse(command string) engine.Command {
	parts := strings.Fields(command)
	if parts[0] == printCommand {
		return (&engine.PrintCommand{Arg: parts[1]})
	} else if parts[0] == addCommand {
		firstNum, err := strconv.Atoi(parts[1])
		if err != nil {
			return (&engine.PrintCommand{Arg: syntaxError + err.Error()})
		}
		secondNum, err := strconv.Atoi(parts[2])
		if err != nil {
			return (&engine.PrintCommand{Arg: syntaxError + err.Error()})
		}
		return (&engine.AddCommand{A: firstNum, B: secondNum})
	} else {
		return (&engine.PrintCommand{Arg: commandError})
	}
}
