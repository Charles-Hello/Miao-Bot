package main

import (
	"github.com/gin-gonic/gin"
	"wechat_http/config"
	"wechat_http/route"
	"wechat_http/util"
)

func main() {
	//读取config.yaml的内容（即可读取端口以及各种的配置）
	config.InitConf()
	//读取系统的参数
	util.GetServerInfo()
	//转入route/gin的配置
	gin.SetMode(gin.ReleaseMode)
	r := route.Setup()
	err := r.Run(":" + config.Config.WechatHttp.Port)
	if err != nil {
		return
	}
}
