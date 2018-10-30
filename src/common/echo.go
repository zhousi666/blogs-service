package common

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// EchoInit echo init
func EchoInit(e *echo.Echo, logpath string, debug bool) {
	e.Debug = debug
	LoggerInit(e, logpath, debug)
	e.HTTPErrorHandler = EchoHTTPErrorHandler(e)
}

// EchoHTTPErrorHandler is a HTTP error handler. It sends a JSON response with status code.
func EchoHTTPErrorHandler(e *echo.Echo) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var (
			code = http.StatusOK
			msg  interface{}
			rmsg string
		)
		errcode := SystemErrorCode
		if he, ok := err.(*echo.HTTPError); ok {
			msg = he.Message
		} else if be, ok := err.(*BizError); ok {
			errcode = be.Code
			msg = be.Msg
		} else {
			msg = err.Error()
		}
		if v, ok := msg.(string); ok {
			rmsg = v
		} else {
			rmsg = fmt.Sprintf("%s", msg)
		}

		if !c.Response().Committed {
			if c.Request().Method == echo.HEAD {
				if err := c.NoContent(code); err != nil {
					goto ERROR
				}
			} else {
				if err := c.JSON(code, ErrorReturns(errcode, rmsg)); err != nil {
					goto ERROR
				}
			}
		}
	ERROR:
		e.Logger.Error(err)
	}
}

// GetAcceptLanguage Get Accept-Language from request header
func GetAcceptLanguage(c echo.Context) string {
	return c.Request().Header.Get("Accept-Language")
}

type SimpleValidator struct {
	Validator *validator.Validate
}

func (cv *SimpleValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
