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
		host, cert, key := common.GetPushHostAndCertFilesPath(set.Type())
		return &ApnClient{
			set:      set,
			host:     host,
			certPath: cert,
			keyPath:  key,
		}
	}
	return &ErrorClient{"service not found"}
}
