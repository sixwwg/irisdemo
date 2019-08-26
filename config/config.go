package config

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"os"
	"ulog_service/iris/core/errors"
)

//服务端配置
type AppConfig struct {
	AppName 	string  `json:"app_name"`
	Port 		string  `json:"port"`
	StaticPath 	string  `json:"static_path"`
 	Mode 		string  `json:"mode"`
}

//初始化服务端配置
func InitConfig() *AppConfig{
	file,err := os.Open("./config.json")
	if err != nil{
		err = errors.New("配置文件打开失败 "+err.Error())
		log4go.Error(err)
	}
	conf := AppConfig{}
	err = json.NewDecoder(file).Decode(&conf)
	if err != nil{
		err = errors.New("配置文件反序列化失败 "+err.Error())
		log4go.Error(err)
	}
	return &conf
}