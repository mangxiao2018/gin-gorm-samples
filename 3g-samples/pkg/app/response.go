package app

import (
	"3g-samples/pkg/e"
	pager "3g-samples/pkg/page"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Page pager.Page  `json:"page"`
}

// (g *Gin)持有这里，不能直接写(c *gin.Context)
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}

// (g *Gin)持有这里，不能直接写(c *gin.Context)
func (g *Gin) ResponseByPage(httpCode, errCode int, data interface{}, page pager.Page) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
		Page: page,
	})
}
