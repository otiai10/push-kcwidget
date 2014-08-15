package model_test

import "github.com/otiai10/push-kcwidget/common"
import "github.com/otiai10/push-kcwidget/model"
import "testing"
import . "github.com/otiai10/mint"
import "time"

func init() {
	common.SetPrefix("test.")
	model.CleanQueue()
	/*
		user := model.User{
			TwitterIdStr: "140021552",
			Name:         "otiai10",
		}
	*/
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
	//
}
