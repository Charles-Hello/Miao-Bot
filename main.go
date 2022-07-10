package main

import (
	"wechat_http/config"
	"wechat_http/route"
	"wechat_http/util"
)

func main() {
	//读取config.yaml的内容（即可读取端口以及各种的配置）
	config.InitConf()
	//初始化cron的任务
	//cron.InitMyCron()
	//读取系统的参数
	util.InitServerInfo()
	//转入route/gin的配置##

	r := route.Setup()
	r.Run(":" + config.Config.WechatHttp.Port)
}
