package plugins

import (
	"fmt"
	"strconv"
	"wechat_http/config"
	"wechat_http/ihttp"
)

func init() {
	ihttp.AddPlugin(func(order ihttp.AddOrder) interface{} {
		order.Cron = ""
		order.Name = "周杰伦你好"
		order.RegStr = "周(.*)伦 ,你(.*)"
		order.Admin = true
		order.RegBool = true
		order.DailyFunction(func() {
			for d, i := range ihttp.AddReg {
				ihttp.PostIHttp(
					ihttp.BuildSendTextMsgBody("第"+strconv.Itoa(d)+"个参数为："+i, order.P.FromWxId))
			}
		})
		return ihttp.FuncFinish(order.Name, order.Cron)
	})

	ihttp.AddPlugin(func(order ihttp.AddOrder) interface{} {
		order.Cron = ""
		order.Name = "我是sad任务"
		order.RegStr = "sad"
		order.RegBool = true
		order.Admin = true
		fmt.Print("sad的测试")
		order.DailyFunction(func() {
			ihttp.PostIHttp(
				ihttp.BuildSendTextMsgBody("我负责sad", order.P.FromWxId))
		})
		return ihttp.FuncFinish(order.Name, order.Cron)
	})

	ihttp.AddPlugin(func(order ihttp.AddOrder) interface{} {
		order.Cron = "*/1 * * * *"
		order.Name = "我是定时任务"
		order.DailyFunction(func() {
			ihttp.PostIHttp(
				ihttp.BuildSendTextMsgBody("我是定时任务", config.Config.MasterWxId))
		})
		return ihttp.FuncFinish(order.Name, order.Cron)
	})

	ihttp.AddPlugin(func(order ihttp.AddOrder) interface{} {
		order.Cron = ""
		order.Name = "指令"
		order.RegStr = "指令"
		order.RegBool = true
		order.Admin = false
		order.DailyFunction(func() {
			switch order.Await(ihttp.Ch, 20, order.P.FromWxId) {
			case "no":
				ihttp.PostIHttp(
					ihttp.BuildSendTextMsgBody("检测到你选择的是n，请输入「结束按钮」",
						order.P.FromWxId))
				switch order.Await(ihttp.Ch, 20, order.P.FromWxId) {
				case "结束":
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody("检测到你选择的是结束\n恭喜你，通关了",
							order.P.FromWxId))
				}
			}
		})
		return ihttp.FuncFinish(order.Name, order.Cron)
	})
}
