package v1

import (
	"github.com/labstack/echo"
)

// RegisterAPI 注册v1版本的API
func RegisterAPI(router *echo.Echo) {
	v1 := router.Group("/v1")

	v1.POST("/sign_up", signUp)

	v1.POST("/sign_in", signIn)
}
