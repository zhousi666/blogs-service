package config

import (
	"github.com/burntsnshi/toml"
)

const (
	CONFIG_PATH = "./config.toml"
)

var (
	Server string
	Config *configStruct
)

//InitConfig 配置初始化: 从 ./config.toml 文件中读取配置，然后赋值给以上定义的全局变量，用于各个模块的初始化。
//除了logpath需要修改，其他都是由DecodeFile自动获取了
func InitConfig() *configStruct {

	config := new(configStruct)
	if _, err := toml.DecodeFile(CONFIG_PATH, config); err != nil {
		panic(err)
	}

	config.Logpath = config.Logpath + "/" + Server + "_server."
	Config = config
	return config
}
