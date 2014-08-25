package controllers

import "github.com/revel/revel"
import "fmt"
import "github.com/otiai10/push-kcwidget/model"

type QueueController struct {
	*revel.Controller
}

func (c *QueueController) Add(
	finish int64,
	id_str,
	message,
	label,
	prefix,
	identifier,
	unit,
	kind,
	missionTitle, //optional
	missionId, //optional
	clientToken string,
) revel.Result {

	if configToken, ok := revel.Config.String("client.token"); ok {
		if configToken != clientToken {
			return c.ErrorOf(fmt.Errorf("Invlid client token"))
		}
	}

	// (1) ユーザ登録の有無を確認する
	user, ok := model.FindUserByTwitterIdStr(id_str)
	if !ok {
		return c.ErrorOf(fmt.Errorf("User not found"))
	}

	// (2) Eventを作る
	// とりあえずこれだけ
	event := model.CreateEventFromRequestParams(
		finish,
		message,
		label,
		prefix,
		identifier,
		unit,
		kind,
		missionId,
		missionTitle,
	)

	// (2) ユーザ情報をメモリでアップデートする
	user = user.SetEvent(event)

	// (3) ユーザ情報をDBにアップデートする
	if e := user.Save(); e != nil {
		revel.ERROR.Println(e)
		return c.ErrorOf(e)
	}

	// (4) エンキューする
	if e := model.Enqueue(finish, user); e != nil {
		revel.ERROR.Println(e)
		return c.ErrorOf(e)
	}
	return c.RenderJson(map[string]interface{}{
		"code":    1000,
		"message": "Enqueue succeeded",
	})
}

func (c *QueueController) ErrorOf(e error) revel.Result {
	return c.RenderJson(map[string]interface{}{
		"code":    2000,
		"error":   "error",
		"message": e.Error(),
	})
}
