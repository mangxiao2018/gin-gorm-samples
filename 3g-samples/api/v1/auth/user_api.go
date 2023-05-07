package auth

import (
	"3g-samples/pkg/app"
	"3g-samples/pkg/e"
	"3g-samples/service/userservices"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserForm struct {
	ID       int64  `form:"id"`
	UserName string `form:"user_name" valid:"Required;MaxSize(50)"`
	NickName string `form:"nick_name" valid:"Required;MaxSize(50)"`
	//UserAvatar string `form:"user_avatar" binding:"required,max=200"`
	//LoginName  string `form:"login_name" binding:"required,max=50"`
	//LoginPwd   string `form:"login_pwd" binding:"required,max=50"`
	//UserNo     string `form:"user_no" binding:"required,max=50"`
	//Gender     int8   `form:"gender" binding:"required"`
	//Email      string `form:"email" binding:"required,max=50"`
	//MobileNo   string `form:"mobile_no" binding:"required,max=20"`
}

// 1.接收来自web端的参数
// 2.对参数进行数据类型转换
// 3.对参数进行校验
// 4.封装response并返回
func AddUser(c *gin.Context) {
	var (
		form UserForm
		appG = app.Gin{C: c}
	)
	// 先绑定
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//~-第一种传参方式和函数调用持有方式----------golang范-------------------------------------
	userservice := userservices.UserService{
		UserName: form.UserName,
		NickName: form.NickName,
	}
	// 查询是否存在，如果存在则提醒并返回
	exists, err := userservice.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USER_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_USER, nil)
		return
	}
	//~-第二种传参方式和函数调用持有方式---------传统方式--------------------------------------
	// 经过一系统校验通过后，正式插入新数据
	// 在上面绑定后使用，顺序不能颠倒
	user := userservices.UserService{
		UserName: form.UserName,
		NickName: form.NickName,
	}
	err = userservices.AddUserService(&user)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Router /user/{id} [put]
func EditUser(c *gin.Context) {
	var (
		form UserForm
		appG = app.Gin{C: c}
	)
	// 先绑定
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	userService := userservices.UserService{
		ID:       form.ID,
		UserName: form.UserName,
		NickName: form.NickName,
	}
	err := userService.EditUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
