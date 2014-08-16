package controllers

import "github.com/revel/revel"

type QueueController struct {
	*revel.Controller
}

func (c *QueueController) Add(finish int64, id_str, message, kind, client_token string) revel.Result {

	revel.INFO.Println(finish, id_str, message, kind, client_token)
	// (1) ユーザ登録の有無を確認する
	// user := model.FindUserByTwitterIdStr(twitterIdStr)
	// if user == nil {
	//     return c.ErrorOf(fmt.Errorf("User not found"))
	// }

	// (2) ユーザ情報をメモリでアップデートする
	// user = model.OverwriteWithRequestParams(params)

	// (3) ユーザ情報をDBにアップデートする
	// if e := user.Save(); e != nil {
	//     return c.ErrorOf(e)

	// (4) エンキューする
	// if e := model.Enqueue(finish, user); e != nil {
	//     return c.ErrorOf(e)
	// }

	return c.RenderJson(map[string]string{
		"hoge": "fuga",
	})
}
