package model

import "push-kcwidget/common"
import "github.com/otiai10/rodeo"

type Enqueuer struct {
	accessor *rodeo.SortedSet
}

type Queue struct {
	Timestamp int64
	User      QueuedUser
}
type QueuedUser struct {
	Name         string
	TwitterIdStr string
}

func Enqueue(timestamp int64, user User) error {
	enqr := &Enqueuer{}
	if e := enqr.InitAccessor(); e != nil {
		return e
	}
	return enqr.Enqueue(timestamp, user)
}

func (enqr *Enqueuer) InitAccessor() (e error) {
	// 毎回インスタンス作るのコスト悪い気がするぞ〜
	host, port := common.GetRedisHostAndPort()
	vaquero, e := rodeo.NewVaquero(host, port)
	if e != nil {
		return e
	}
	accessor, e := vaquero.Tame(common.Prefix()+"queues", Queue{})
	if e != nil {
		return e
	}
	enqr.accessor = accessor
	return
}

func (enqr *Enqueuer) Enqueue(timestamp int64, user User) (e error) {
	// if accessor is not initialized return error
	return enqr.accessor.Add(timestamp, Queue{Timestamp: timestamp, User: QueuedUser{user.Name, user.TwitterIdStr}})
}

/**
 * for dev
 */
func CleanQueue() (e error) {
	host, port := common.GetRedisHostAndPort()
	vaquero, e := rodeo.NewVaquero(host, port)
	if e != nil {
		return
	}
	return vaquero.Delete(common.Prefix() + "queues")
}
