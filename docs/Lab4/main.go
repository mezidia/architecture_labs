package main

import (
	"time"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type commandsQueue struct {
	c []Command
}

func (q *commandsQueue) push(cmd Command) {
	q.c = append(q.c, cmd)
}

func (q *commandsQueue) pull() Command {
	res := q.c[0]
	q.c[0] = nil
	q.c = q.c[1:]
	return res
}

type Loop struct {
	queue *commandsQueue

	stopRequest bool
	stopConfirm chan struct{}
}

func (l *Loop) Post(cmd Command) {
	l.queue.push(cmd)
}

func (l *Loop) Start() {
	l.queue = &commandsQueue{}
	l.stopConfirm = make(chan struct{})
	go func() {
		for {
			// TODO: Respect stopRequest and message queue empty state.
			cmd := l.queue.pull()
			cmd.Execute(l)
		}
		l.stopConfirm <- struct{}{}
	}()
}

func (l *Loop) AwaitFinish() {
	l.stopRequest = true
	<-l.stopConfirm
}

func main() {

	loop := new(Loop)

	loop.Start()

	loop.Post(&printCommand{arg: "hello"})

	loop.Post(&addCommand{a: 4, b: 5})

	time.Sleep(1 * time.Second)

	loop.AwaitFinish()

}
