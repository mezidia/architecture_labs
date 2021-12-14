package engine

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
	queue       *commandsQueue
	stopConfirm chan bool

	stopRequest  bool
	pauseRequest bool
}

func (l *Loop) Post(cmd Command) {
	l.queue.push(cmd)
	if l.pauseRequest && !l.stopRequest {
		l.Routine()
	}
}

func (l *Loop) Routine() {
	l.pauseRequest = false
	go func() {
		for {
			if len(l.queue.c) > 0 {
				cmd := l.queue.pull()
				cmd.Execute(l)
			} else if l.stopRequest {
				l.stopConfirm <- true
				return
			} else {
				l.pauseRequest = true
				return
			}
		}
	}()
}

func (l *Loop) Start() {
	l.queue = &commandsQueue{}
	l.stopConfirm = make(chan bool, 1)
	l.Routine()
}

func (l *Loop) AwaitFinish() {
	l.stopRequest = true
	<-l.stopConfirm
}
