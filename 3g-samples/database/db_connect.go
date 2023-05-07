package database

import (
	"3g-samples/pkg/logging"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var _db *gorm.DB

func InitDatabase() {
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	dbname := viper.GetString("mysql.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("数据库连接失败, error=" + err.Error())
	}
	db, err := _db.DB()
	if err != nil {
		logging.Fatal("db connected error", err)
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
}

func GetDB() *gorm.DB {
	return _db
}

/*通过session方式获取会话连接*/
func GetSession() *gorm.DB {
	return _db.Session(&gorm.Session{PrepareStmt: true})
}
