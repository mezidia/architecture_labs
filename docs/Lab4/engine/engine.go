package engine

import (
	"fmt"
	"strconv"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type printCommand struct {
	arg string
}

func (pc *printCommand) Execute(handler Handler) {
	fmt.Println(pc.arg)
}

type addCommand struct {
	a, b int
}

func (ac *addCommand) Execute(handler Handler) {
	res := ac.a + ac.b
	handler.Post(&printCommand{arg: strconv.Itoa(res)})
}
