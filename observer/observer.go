package observer

import "time"
import "fmt"

type Observer struct {
	closer chan error
	dead   bool
	err    error
}

func New() *Observer {
	return &Observer{
		closer: make(chan error),
	}
}

func (o *Observer) Start() chan error {
	go o.run()
	return o.closer
}

func (o *Observer) Close(message string) *Observer {
	o.dead = true
	o.closer <- fmt.Errorf(message)
	return o
}

func (o *Observer) run() {
	for {
		select {
		case now := <-time.Tick(1 * time.Second):
			fmt.Println(now)
		case err := <-o.closer:
			o.err = err
			break
		}
	}
	return
}
