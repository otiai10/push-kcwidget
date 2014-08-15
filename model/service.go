package model

import "github.com/otiai10/push-kcwidget/common"

type Service struct {
	Type  common.PushType
	Token string
}

var nameTypeMapping = map[string]common.PushType{
	"apn": common.PushTypeApple,
}

func GetPushTypeByName(serviceName string) common.PushType {
	if t, ok := nameTypeMapping[serviceName]; ok {
		return t
	}
	return common.PushTypeApple
}
