package model

import "github.com/otiai10/push-kcwidget/common"

// これモデルじゃね？
type PushSet struct {
	typ    common.PushType
	token  string
	events []Event
}

func NewPushSet(typ common.PushType, token string, events []Event) PushSet {
	return PushSet{
		typ:    typ,
		token:  token,
		events: events,
	}
}
func (set PushSet) Type() common.PushType {
	return set.typ
}
func (set PushSet) Token() string {
	return set.token
}
func (set PushSet) Events() []Event {
	return set.events
}
