package service

import "push-kcwidget/common"

type PushClient interface {
	Send() error
}

type PushSet interface {
	Type() common.PushType
	Token() string
	Events() []Event
	// GetParams()?
}

type Event interface{}

func NewClient(set PushSet) PushClient {
	switch set.Type() {
	case common.PushTypeApple:
		return &ApnClient{set: set}
	}
	return &ErrorClient{"service not found"}
}
