package service

import "github.com/otiai10/push-kcwidget/common"
import "github.com/otiai10/push-kcwidget/model"

type PushClient interface {
	Send() error
}

type PushSet interface {
	Type() common.PushType
	Token() string
	Events() []model.Event
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
