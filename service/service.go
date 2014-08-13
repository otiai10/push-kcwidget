package service

import "github.com/otiai10/push-kcwidget/common"

type PushClient interface {
	Send() error
}

type PushSet interface {
	Type() common.PushType
	Token() string
	Events() []Event
}

type Event interface {
	// ToParams()?
}

func NewClient(set PushSet) PushClient {
	switch set.Type() {
	case common.PushTypeApple:
		return &ApnClient{set: set}
	}
	return &ErrorClient{"service not found"}
}
