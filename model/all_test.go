package model_test

import "github.com/otiai10/push-kcwidget/common"
import "github.com/otiai10/push-kcwidget/model"
import "testing"
import . "github.com/otiai10/mint"
import "time"

func init() {
	common.SetPrefix("test.")
	model.CleanQueue()
	user := model.User{
		TwitterIdStr: "140021552",
		Name:         "otiai10",
	}
	user.Save()
}

func TestUser(t *testing.T) {
	user := model.User{}
	Expect(t, user).TypeOf("model.User")
}

func TestEnqueue(t *testing.T) {
	user := model.User{
		TwitterIdStr: "140021552",
		Name:         "otiai10",
	}
	e := model.Enqueue(time.Now().Unix(), user)
	Expect(t, e).ToBe(nil)
}

func TestFindUserByTwitterIdStr(t *testing.T) {
	user, ok := model.FindUserByTwitterIdStr("140021552")
	Expect(t, ok).ToBe(true)
	Expect(t, user.Name).ToBe("otiai10")

	user, ok = model.FindUserByTwitterIdStr("000")
	Expect(t, ok).ToBe(false)
	Expect(t, user.Name).ToBe("")
}

func TestCreateEventFromRequestParams(t *testing.T) {
	event := model.CreateEventFromRequestParams(
		10000000000000,
		"This is message to be push-notified",
		"遠征帰投",
		"第",
		"3",
		"艦隊",
		"mission-finish",
		"38", "東京急行(弐)", "and", "any", "other", "optional", "strings",
	)
	Expect(t, event).TypeOf("model.Event")

	Expect(t, len(event.OptionalStrings)).ToBe(7)
}
