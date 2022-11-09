package robot

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-defaults"
	"net/http"
	"time"
	"wechat_http/ihttp"
	"wechat_http/plugins"
	"wechat_http/telegram_bot"
)

func Robot(c *gin.Context) {
	//每次执行都要把telegram_bot.Return给重置了
	telegram_bot.Return.OpSlice = nil
	defaults.SetDefaults(telegram_bot.Return)

	var AddOrder ihttp.AddOrder
	err := c.BindJSON(&AddOrder.P)
	if err != nil {
		return
	}
	ihttp.Con = true
	go func() {
		ihttp.Ch <- AddOrder.P
	}()

	now := time.Now().UnixNano()
	jsonBite, _ := json.MarshalIndent(AddOrder.P, "", "    ")
	jsonStr := string(jsonBite)
	fmt.Println(jsonStr)

	plugins.TraverseFunction(AddOrder) //这里开始遍历切片函数，使得装载函数全部执行
	after := time.Now().UnixNano()
	cost := after - now
	fmt.Println("golang消耗时间=[", cost, "]纳秒")

	//用变量来接收
	data := gin.H{"Response": telegram_bot.Return}

	c.JSON(http.StatusOK, data) //后续跟telegram人形交接

}
