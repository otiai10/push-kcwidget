package service_test

import "github.com/otiai10/push-kcwidget/service"
import "github.com/otiai10/push-kcwidget/common"
import "testing"
import . "github.com/otiai10/mint"

type DummySet struct {
	typ common.PushType
}

func (set DummySet) Type() common.PushType {
	return set.typ
}
func (set DummySet) Events() (events []service.Event) {
	return
}
func (set DummySet) Token() string {
	return "xxxxxxxx xxxxxxxx xxxxxxxx xxxxxxxx"
}

func TestNewClient(t *testing.T) {
	set := DummySet{common.PushTypeApple}
	client := service.NewClient(set)
	Expect(t, client).TypeOf("*service.ApnClient")
}
