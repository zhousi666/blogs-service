package common

import "github.com/labstack/echo"

// PageBody 分页结果
type PageBody struct {
	Current int `json:"current"`
	Total   int `json:"total,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// ReturnBody 返回值封装
type ReturnBody struct {
	Errcode string      `json:"errcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Page    PageBody    `json:"page"`
}

// PageParams 分页参数，用于mixin到请求对象中
type PageParams struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page" validate:"gte=1"`
}

// JSONReturns API返回值的统一封装，直接做json返回。
// `data`为需要返回的数据
// `pages`为翻页数据，不是必须要有。顺序为: page, total, per_page，其中per_page如果不设置则默认为20。
// 如果使用了这个参数，则 page, total必须有
func JSONReturns(c echo.Context, data interface{}, pages ...int) error {
	var page PageBody
	if len(pages) > 0 {
		current := pages[0]
		total := pages[1]
		perPage := 20
		if len(pages) > 2 {
			perPage = pages[2]
		}
		page = PageBody{
			Current: current,
			Total:   total,
			PerPage: perPage,
		}
	}
	returns := &ReturnBody{
		Errcode: OkCode,
		Data:    data,
		Page:    page,
	}

	return c.JSON(200, returns)
}

// ErrorReturns 发生错误的时候的返回值封装
func ErrorReturns(errcode string, msg string) *ReturnBody {
	return &ReturnBody{
		Errcode: errcode,
		Msg:     msg,
		Page:    PageBody{},
	}
}
