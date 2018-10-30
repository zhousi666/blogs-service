package common

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/color"
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

//按照 github.com/labstack/echo/middleware/recover.go 来重写
type (
	RecoverConfig struct {
		Skipper           middleware.Skipper
		StackSize         int  `json:"stack_size"`
		DisableStackAll   bool `json:"disable_stack_all"`
		DisablePrintStack bool `json:"disable_print_stack"`
	}
)

var (
	DefaultRecoverConfig = RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10,
		DisableStackAll:   false,
		DisablePrintStack: false,
	}
)

func Recover() echo.MiddlewareFunc {
	return RecoverWithConfig(DefaultRecoverConfig)
}

func RecoverWithConfig(config RecoverConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultRecoverConfig.Skipper
	}
	if config.StackSize == 0 {
		config.StackSize = DefaultRecoverConfig.StackSize
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			defer func() {
				if r := recover(); r != nil {
					var err error
					switch r := r.(type) {
					case error:
						err = r
					default:
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, config.StackSize)
					length := runtime.Stack(stack, !config.DisableStackAll)
					if !config.DisablePrintStack {
						Logger.Debugf("[%s] %s %s\n", color.Red("PANIC RECOVER"), err, stack[:length])
					}
					bizerr := NewBizError(SystemErrorCode, "", string(stack[:length]))
					c.Set(ContextError, bizerr)
					c.Error(bizerr)
				}
			}()
			return next(c)
		}
	}
}

type SimpleValidator struct {
	Validator *validator.Validate
}

func (cv *SimpleValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
