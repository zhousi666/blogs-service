package v1

import (
	"blogs-service/common"
	"blogs-service/service/blogs"

	"github.com/labstack/echo"
)

func signUp(c echo.Context) error {
	params := new(SignUpStruct)
	if err := c.Bind(params); err != nil {
		return err
	}
	err := blogs.SignUp(params.UserName, params.Password, params.Sex, params.TelNum)
	if err != nil {
		common.Logger.Error(err)
		return err
	}
	ret := "sign up ok!"
	return common.JSONReturns(c, ret)
}

func signIn(c echo.Context) error {
	params := new(SignInStruct)
	if err := c.Bind(params); err != nil {
		return err
	}
	err := blogs.SignIn(params.UserName, params.Password)
	if err != nil {
		common.Logger.Error(err)
		return err
	}
	ret := "Wecome come back, " + params.UserName + " !"
	return common.JSONReturns(c, ret)
}
