package controllers

import "fmt"
import "github.com/revel/revel"
import "github.com/otiai10/push-kcwidget/model"

type UserController struct {
	*revel.Controller
}

type UserRegistrationParams struct{}

func (params UserRegistrationParams) ToMap() (m map[string]string) {
	return
}

func (c *UserController) Register(username, idStr, deviceToken, service, clientToken string) revel.Result {

	if configToken, ok := revel.Config.String("client.token"); ok {
		if configToken != clientToken {
			return c.ErrorOf(fmt.Errorf("Invlid client token"))
		}
	}

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

	return c.RenderJson(map[string]interface{}{
		"code":    1000,
		"message": "User registration succeeded",
	})
}

func (c *UserController) Get(twitterIdStr, clientToken string) revel.Result {

	if configToken, ok := revel.Config.String("client.token"); ok {
		if configToken != clientToken {
			return c.ErrorOf(fmt.Errorf("Invlid client token"))
		}
	}

	user, ok := model.FindUserByTwitterIdStr(twitterIdStr)
	return c.RenderJson(map[string]interface{}{
		"code":    1000,
		"message": ok,
		"user":    user.FilterPrivateInfo().SortEvents(),
	})
}

func (c *UserController) ErrorOf(e error) revel.Result {
	return c.RenderJson(map[string]interface{}{
		"code":    2000,
		"message": e.Error(),
	})
}
