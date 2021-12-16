package engine

import "sync"

type commandsQueue struct {
	sync.Mutex
	c []Command
}

func (q *commandsQueue) push(cmd Command) {
	q.Lock()
	q.c = append(q.c, cmd)
	q.Unlock()
}

func (q *commandsQueue) pull() Command {
	q.Lock()
	res := q.c[0]
	q.c[0] = nil
	q.c = q.c[1:]
	q.Unlock()
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
			var queueLength = len(l.queue.c)
			if queueLength > 0 {
				cmd := l.queue.pull()
				cmd.Execute(l)
			} else if !l.stopRequest {
				l.pauseRequest = true
				return
			} else {
				l.stopConfirm <- true
				return
			}
		}
	}()
}

func (l *Loop) Start() {
	l.queue = &commandsQueue{}
	l.stopConfirm = make(chan bool)
	l.Routine()
}

func (l *Loop) AwaitFinish() {
	l.stopRequest = true
	<-l.stopConfirm
}
