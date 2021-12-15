package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/mezidia/architecture_labs/tree/main/docs/Lab4/engine"
)

var (
	inputFile      = "examples.txt"
	syntaxError    = "SYNTAX ERROR: "
	commandError   = "unknown command"
	emptyLineError = "empty line"
	printCommand   = "print"
	addCommand     = "add"
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
}

func parse(command string) engine.Command {
	if len(command) <= 0 {
		return (&engine.PrintCommand{Arg: emptyLineError})
	}

	parts := strings.Fields(command)

	if parts[0] == printCommand {
		return (&engine.PrintCommand{Arg: parts[1]})
	} else if parts[0] == addCommand {
		firstNum, firstNumErr := strconv.Atoi(parts[1])
		if firstNumErr != nil {
			return (&engine.PrintCommand{Arg: syntaxError + firstNumErr.Error()})
		}
		secondNum, secondNumErr := strconv.Atoi(parts[2])
		if secondNumErr != nil {
			return (&engine.PrintCommand{Arg: syntaxError + secondNumErr.Error()})
		}
		return (&engine.AddCommand{A: firstNum, B: secondNum})
	} else {
		return (&engine.PrintCommand{Arg: commandError})
	}
}
