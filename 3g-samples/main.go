package main

import (
	"3g-samples/database"
	"3g-samples/pkg/logging"
	"3g-samples/pkg/setting"
	"3g-samples/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	setting.InitConfigs("dev")
	database.InitDatabase()
	logging.InitLogger()
}

func main() {
	r := gin.Default()
	router.InitRouter(r)
	port := fmt.Sprintf(":%s", viper.GetString("server.port"))
	r.Run(port) //default:0.0.0.0:8080
}
