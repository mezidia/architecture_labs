package main

import (
	"time"

	"github.com/mezidia/architecture_labs/tree/main/docs/Lab4/engine"
)

func main() {

	loop := new(engine.Loop)

	loop.Start()

	loop.Post(&printCommand{arg: "hello"})

	loop.Post(&addCommand{a: 4, b: 5})

	time.Sleep(1 * time.Second)

	loop.AwaitFinish()

}
