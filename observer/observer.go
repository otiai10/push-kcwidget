package observer

type Observer struct{}

func New() *Observer {
	return &Observer{}
}

func (o *Observer) Start() chan error {
	ch := make(chan error)
	return ch
}
