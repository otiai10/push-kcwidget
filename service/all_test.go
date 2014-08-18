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
	return "324175ddc7f8ba3944f60cb88b30b955b6d215570e59b5088a531580742b66e6"
}
func init() {
	common.SetPrefix("test.")
}

func TestNewClient(t *testing.T) {
	set := DummySet{common.PushTypeApple}
	client := service.NewClient(set)
	Expect(t, client).TypeOf("*service.ApnClient")
}

/* It works, see common.GetPushHostAndCertFilesPath
func TestApnClient_Send(t *testing.T) {
	set := DummySet{common.PushTypeApple}
	client := service.NewClient(set)
	e := client.Send()
	Expect(t, e).ToBe(nil)
}
 */
