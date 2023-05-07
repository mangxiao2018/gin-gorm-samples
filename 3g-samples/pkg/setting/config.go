package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfigs(env string) {
	workRootPath, _ := os.Getwd()
	defaultConfigFileName := fmt.Sprintf("app.%s", env)
	viper.SetConfigName(defaultConfigFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workRootPath + "/configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Get yamFile error: %s", err))
	}

}
