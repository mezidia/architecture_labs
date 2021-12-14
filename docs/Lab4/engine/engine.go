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

type PrintCommand struct {
	Arg string
}

func (pc *PrintCommand) Execute(handler Handler) {
	fmt.Println(pc.Arg)
}

type AddCommand struct {
	A, B int
}

func (ac *AddCommand) Execute(handler Handler) {
	res := ac.A + ac.B
	handler.Post(&PrintCommand{Arg: strconv.Itoa(res)})
}
