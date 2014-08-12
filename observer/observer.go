package observer

import "time"
import "fmt"
import "github.com/otiai10/rodeo"
import "github.com/otiai10/push-kcwidget/common"
import "github.com/otiai10/push-kcwidget/model"

type Observer struct {
	closer   chan error
	dead     bool
	err      error
	accessor *rodeo.SortedSet
}

var iteration = 1 * time.Second

func initAccessor() (ss *rodeo.SortedSet, e error) {
	host, port := common.GetRedisHostAndPort()
	vaq, e := rodeo.NewVaquero(host, port)
	ss, e = vaq.Tame(common.Prefix()+"queues", model.Queue{})
	return
}

func New() *Observer {
	accessor, _ := initAccessor()
	return &Observer{
		closer:   make(chan error),
		accessor: accessor,
	}
}

func (o *Observer) Start() chan error {
	go o.run()
	return o.closer
}

func (o *Observer) Kill(message string) *Observer {
	o.dead = true
	o.closer <- fmt.Errorf(message)
	return o
}

func (o *Observer) run() {
	for {
		select {
		case now := <-time.Tick(iteration):
			o.execute(now)
		case err := <-o.closer:
			o.err = err
			break
		}
	}
	return
}

func (o *Observer) execute(now time.Time) {
	// 現時点までをスコアに持つScoredValueを全て取得
	queues := o.accessor.Find(0, now.Unix())
	if len(queues) < 1 {
		return
	}
	for _, q := range queues {
		queue := q.Retrieve().(*model.Queue)
		fmt.Printf("%+v\n", queue)
		if e := o.callPushServiceFromQueue(queue); e == nil {
			o.accessor.Remove(queue)
		}
	}
}

func (o *Observer) callPushServiceFromQueue(queue *model.Queue) (e error) {
	// Do something
	return
}
