package controllers

import "github.com/revel/revel"

type UserController struct {
	*revel.Controller
}

type UserRegistrationParams struct{}

func (params UserRegistrationParams) ToMap() (m map[string]string) {
	return
}

func (c *UserController) Register(params UserRegistrationParams) revel.Result {

	// (1) ユーザを取得する
	// old := model.FindUserByTwitterIdStr(twitterIdStr)

	// (2) 新規ユーザを作る
	// user := model.CreaetUserWithRegisterParams(params)

	// (3) マージする
	// user = model.MergeUser(user, old)

	// (4) ユーザ情報をDBにアップデートする
	// if e := user.Save(); e != nil {
	//     return c.ErrorOf(e)

	return c.RenderJson(map[string]string{
		"message": "User registration succeeded",
	})
}
