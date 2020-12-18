package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	conf := viper.New()
	env := os.Getenv("GO_ENV")
	if env == "" {
		panic("GO_ENV is not found!")
	}
	conf.SetConfigName("config_" + env) // 设置配置文件名 (不带后缀)

	conf.AddConfigPath("../configs") // 第一个搜索路径
	conf.SetConfigType("yml")        // 文件类型
	err := conf.ReadInConfig()       // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return conf
}
