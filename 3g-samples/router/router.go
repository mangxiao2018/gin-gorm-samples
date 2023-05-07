package router

import (
	"3g-samples/api/v1/auth"
	"3g-samples/api/v1/orderapi"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("/adduser", auth.AddUser)                        // 新增用户
	r.PUT("/user/:id", auth.EditUser)                       // 编码用户
	r.GET("/queryOrdersByPage", orderapi.QueryOrdersByPage) // 查询订单:多表分页不定条件查询
	r.POST("/addOrder", orderapi.AddOrder)                  // 新增订单、订单明细:Json数据使用事务多表写入
}
