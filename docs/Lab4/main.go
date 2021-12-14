package main

import (
	"time"
)

func main() {

	loop := new(Loop)

	loop.Start()

	loop.Post(&printCommand{arg: "hello"})

	loop.Post(&addCommand{a: 4, b: 5})

	time.Sleep(1 * time.Second)

	loop.AwaitFinish()

}
