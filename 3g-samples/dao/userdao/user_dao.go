package userdao

import (
	"3g-samples/database"
	"3g-samples/pkg/logging"
	"3g-samples/pkg/util"
	"gorm.io/gorm"
	"time"
)

// CreateTime,UpdateTime设置默认值是必须加上NOT NULL否则不起作用
type User struct {
	ID         int64     `gorm:"primary_key:yes"`
	UserName   string    `gorm:"column:user_name"`
	NickName   string    `gorm:"column:nick_name"`
	UserAvatar string    `gorm:"column:user_avatar"`
	LoginName  string    `gorm:"column:login_name"`
	LoginPwd   string    `gorm:"column:login_pwd"`
	UserNo     string    `gorm:"column:user_no"`
	Gender     int8      `gorm:"column:gender"`
	Email      string    `gorm:"column:email"`
	MobileNo   string    `gorm:"column:mobile_no"`
	CreateAt   time.Time `gorm:"column:create_at;NOT NULL;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:update_at;NOT NULL;autoUpdateTime"`
	CreateUser int64     `gorm:"column:create_user;NOT NULL;default:1"`
	UpdateUser int64     `gorm:"column:update_user;NOT NULL;default:1"`
	Yn         int       `gorm:"column:yn;NOT NULL;default:1"`
}

// 1、操作数据库或操作redis:增删改查、分页查、批量插入、批量删除等
func AddUserBySingle(u *User) error {
	err := database.GetDB().Create(u).Error
	if err != nil {
		logging.Error("新增一条用户信息失败:", err)
		return err
	}
	return nil
}

// 根据用户名判断是否在库表中已存在，用于新增一条用户数据前的验证
func ExistByName(userName string) (bool, error) {
	var user User
	err := database.GetDB().Select("id").Where("user_name = ? AND yn = ?", userName, 1).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// 编辑后更新指定用户数据
func EditUser(u *User) error {
	if err := database.GetDB().Model(&User{}).Where("id = ? AND yn = ?", u.ID, 1).Updates(u).Error; err != nil {
		return err
	}
	return nil
}

// 根据用户编号查询该用户全量信息
func (user *User) QueryUserByNo(tx *gorm.DB) (bool, error) {
	var u User
	err := tx.Where("user_no = ? AND yn = ?", user.UserNo, 1).Find(&u).Error
	util.StructUtils(&u, &user)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if u.ID > 0 {
		return true, nil
	}
	return false, nil
}
