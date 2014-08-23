package controllers

import "github.com/revel/revel"
import "github.com/otiai10/push-kcwidget/model"

type UserController struct {
	*revel.Controller
}

type UserRegistrationParams struct{}

func (params UserRegistrationParams) ToMap() (m map[string]string) {
	return
}

func (c *UserController) Register(username, idStr, deviceToken, service string) revel.Result {

	revel.INFO.Println(username, idStr, deviceToken)

	if service == "" {
		service = "apn"
	}

	user := model.CreaetOrMergeUserWithRegisterParams(
		username,
		idStr,
		deviceToken,
		service,
	)
	if e := user.Save(); e != nil {
		return c.ErrorOf(e)
	}

	return c.RenderJson(map[string]string{
		"message": "User registration succeeded",
	})
}

func (c *UserController) Get(twitterIdStr string) revel.Result {
	user, ok := model.FindUserByTwitterIdStr(twitterIdStr)
	return c.RenderJson(map[string]interface{}{
		"message": ok,
		"user":    user.FilterPrivateInfo(),
	})
}

func (c *UserController) ErrorOf(e error) revel.Result {
	return c.RenderJson(map[string]string{
		"message": e.Error(),
	})
}
