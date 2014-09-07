package observer

import "time"
import "fmt"
import "github.com/otiai10/rodeo"
import "github.com/otiai10/push-kcwidget/common"
import "github.com/otiai10/push-kcwidget/model"
import "github.com/otiai10/push-kcwidget/service"
import "github.com/revel/revel" // to log

type Observer struct {
	closer   chan error
	dead     bool
	err      error
	accessor *rodeo.SortedSet
}

var iteration = 20 * time.Second
var onFail = make(chan model.User)

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
			return
		case userSnapShot := <-onFail:
			revel.ERROR.Printf("[To    USER] %+v", userSnapShot)
			// something to handle failure
		}
	}
	return
}

func (o *Observer) fail(userSnapShot model.User) {
	onFail <- userSnapShot
}

func (o *Observer) execute(now time.Time) {
	// 現時点までをスコアに持つScoredValueを全て取得
	queues := o.accessor.Find(0, now.Unix())
	if len(queues) < 1 {
		return
	}
	for _, q := range queues {
		queue := q.Retrieve().(*model.Queue)
		if e := o.callPushServiceFromQueue(queue); e == nil {
			o.accessor.Remove(queue)
		}
	}
}

func (o *Observer) callPushServiceFromQueue(queue *model.Queue) (e error) {
	go o.sendNotificationAndCleanUpUserEvent(queue)
	return
}
func (o *Observer) createPushSets(queue *model.Queue) (sets []model.PushSet, user model.User) {
	// queueからuserを取得
	// userが見つからない的なやつは、queue成功でよいのでスルー
	user, _ = model.FindUserByTwitterIdStr(queue.User.TwitterIdStr)
	events := user.FindReadyEvents()
	for _, s := range user.Services {
		sets = append(sets, model.NewPushSet(s.Type, s.Token, events))
	}
	return
}
func (o *Observer) sendNotificationAndCleanUpUserEvent(queue *model.Queue) {
	checked := time.Unix(queue.Timestamp, 0)
	sets, user := o.createPushSets(queue)
	var e error
	for _, set := range sets {
		// queueはあるがUser.Eventsが更新され
		// ReadyEventsは無い場合ことは十分ある
		if len(set.Events()) == 0 {
			continue
		}
		client := service.NewClient(set)
		e = client.Send()
	}
	if e != nil {
		revel.ERROR.Printf("[PUSH ERROR] %+v", e)
		o.fail(user)
		return
	}
	user.CleanUpEvents(checked)
}
