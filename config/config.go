package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Config 保存程序的所有配置信息
var Config = new(AppConfig)

type AppConfig struct {
	*WechatHttp `yaml:"wechatHttp"`
	*IHttp      `yaml:"iHttp"`
	*Redis      `yaml:"redis"`
}

type WechatHttp struct {
	Token string `yaml:"token"`
	Port  string `yaml:"port"`
}

type IHttp struct {
	URL           string `yaml:"url"`
	RobotWxId     string `yaml:"robotWxId"`
	MasterWxId    string `yaml:"masterWxId"`
	Authorization string `yaml:"authorization"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Pwd  string `yaml:"pwd"`
	Db   int    `yaml:"db"`
}

func InitConf() {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := viper.New()
	config.AddConfigPath(path)     //设置读取的文件路径
	config.SetConfigName("config") //设置读取的文件名
	config.SetConfigType("yml")    //设置文件的类型
	err = config.ReadInConfig()
	if err != nil {
		// 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}

	// 配置信息绑定到结构体变量
	err = config.Unmarshal(Config)
	if err != nil {
		fmt.Printf("viper.Unmarshal() failed, err:%v\n", err)
	}
}
