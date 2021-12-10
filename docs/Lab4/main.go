package main

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
}

func (l *Loop) Post(cmd Command) {
	l.queue.push(cmd)
}

func (l *Loop) Start() {
	l.queue = &commandsQueue{}
	go func() {
		for {
			cmd := l.queue.pull()
			cmd.Execute(l)
		}
	}()
}

func (l *Loop) AwaitFinish() {

}

func main() {

	loop := new(Loop)

	loop.Start()

	loop.Post(&printCommand{arg: "hello"})

	loop.AwaitFinish()

}
