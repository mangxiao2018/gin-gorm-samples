package userservices

import (
	"3g-samples/dao/userdao"
)

type UserService struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
}

// 1、业务逻辑编写
// 1.1、分页
// 1.2、验证db或redis里中是否存在该值
// 1.3、文件操作类的需求，在这里进行文件打开、处理文件内容、文件关闭等操作
// 1.3、业务逻辑
// 2、返回时可以多返回值，这样在api层就可以进行判断和封装response
func AddUserService(u *UserService) error {
	daoUser := userdao.User{
		UserName: u.UserName,
		NickName: u.NickName,
	}
	return userdao.AddUserBySingle(&daoUser)
}

// 此函数的传参方式、调用方式与AddUserService完全不一样，此方式更golang范
func (user *UserService) ExistByName() (bool, error) {
	return userdao.ExistByName(user.UserName)
}

// 编辑修改指定用户信息
func (user *UserService) EditUser() error {
	daoUser := userdao.User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
	}
	return userdao.EditUser(&daoUser)
}
