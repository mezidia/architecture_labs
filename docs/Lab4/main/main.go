package main

import (
	"github.com/mezidia/architecture_labs/tree/main/docs/Lab4/engine"
)

func main() {

	loop := new(engine.Loop)

	loop.Start()
	loop.Post(&engine.PrintCommand{Arg: "hello"})
	loop.Post(&engine.AddCommand{A: -20, B: 6})

	loop.AwaitFinish()
	loop.Post(&engine.PrintCommand{Arg: "Should end!"})
}
