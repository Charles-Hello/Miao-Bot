package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat_http/api"
	"wechat_http/config"
	"wechat_http/response"
	"wechat_http/robot"
)

func Setup() *gin.Engine {
	r := gin.Default()
	// 开启可爱猫的对接
	r.POST("/robot", robot.Robot)

	// 对外提供api
	group := r.Group("/api")
	group.Use(middleWareCheckToken())
	{
		// 发送文本消息给机器人主人
		group.GET("/sendTextMsg", api.SendTextMsg)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, response.ErrorResponse("你想干嘛"))
	})
	return r
}

// middleWareCheckToken 校验token
func middleWareCheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusOK, response.ErrorResponse("token不能为空"))
			c.Abort()
			return
		}
		if token != config.Config.WechatHttp.Token {
			c.JSON(http.StatusOK, response.ErrorResponse("token错误"))
			c.Abort()
			return
		}
	}
}
