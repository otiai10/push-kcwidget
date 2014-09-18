package service

import "fmt"
import "strings"
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

func getMessage(events []model.Event) string {
	switch len(events) {
	case 0:
		return "エラーメッセージ"
	case 1:
		return events[0].Message
	default:
		return squashMessages(events)
	}
}
func squashMessages(events []model.Event) string {
	message := events[0].Message + "\n"
	counts := map[string]int{}
	for _, ev := range events[1:] {
		counts[ev.Label]++
	}
	pool := []string{"(他"}
	for label, count := range counts {
		pool = append(pool, fmt.Sprintf("%v%d", label, count))
	}
	pool = append(pool, ")")
	return message + strings.Join(pool, " ")
}
