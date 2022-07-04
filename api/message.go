package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat_http/response"
)

func SendTextMsg(c *gin.Context) {
	msg := c.Query("msg")
	if msg == "" {
		c.JSON(http.StatusOK, response.ErrorResponse("缺少msg参数"))
		return
	}
	//c.JSON(http.StatusOK, response.SuccessResponse(service.SendMsgToMe(msg)))
}
