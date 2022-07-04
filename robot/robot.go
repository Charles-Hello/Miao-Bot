package robot

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wechat_http/ihttp"
	"wechat_http/plugins"
)

func Robot(c *gin.Context) {
	var AddOrder ihttp.AddOrder
	c.BindJSON(&AddOrder.P)
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
	data := gin.H{"success": false}
	c.JSON(http.StatusOK, data) //后续跟telegram人形交接。
}
