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
    missionTitle,//optional
    missionId,//optional
    client_token string,
    ) revel.Result {

	revel.INFO.Println(
        finish,
        id_str,
        message,
        label,
        prefix,
        identifier,
        unit,
        kind,
        client_token,
        missionTitle,
        missionId,
    )
	// (1) ユーザ登録の有無を確認する
	user, ok := model.FindUserByTwitterIdStr(id_str)
	if ! ok {
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
    revel.INFO.Printf("%+v\n", user)
    revel.INFO.Printf("%+v\n", event)

	// (2) ユーザ情報をメモリでアップデートする
	user = user.SetEvent(event)

	// (3) ユーザ情報をDBにアップデートする
	if e := user.Save(); e != nil {
	     return c.ErrorOf(e)
    }

	// (4) エンキューする
	if e := model.Enqueue(finish, user); e != nil {
	     return c.ErrorOf(e)
	}
	return c.RenderJson(map[string]string{
		"message": "Enqueue succeeded",
	})
}

func (c *QueueController) ErrorOf(e error) revel.Result {
    return c.RenderJson(map[string]string{
        "error": "error",
        "message": e.Error(),
    })
}
