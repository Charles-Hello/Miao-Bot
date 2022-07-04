package plugins

import (
	"fmt"
	"wechat_http/ihttp"
)

func init() {
	ihttp.AddPlugin(func(order ihttp.AddOrder) interface{} {
		order.Cron = ""
		order.Name = "激活插件"
		order.RegStr = "^激活(.*)"
		order.Admin = true
		order.RegBool = true
		order.DailyFunction(func() {
			for j, i := range ihttp.PluginNameList {
				if i == ihttp.AddReg[0] {
					ihttp.PluginQueue = append(ihttp.PluginQueue, ihttp.PluginList[j])
					//ihttp.PluginQueue = append(ihttp.PluginQueue[:j], append([]ihttp.PluginFunc{ihttp.PluginList[j-1]}, ihttp.PluginQueue[j:]...)...)
					//fmt.Println(ihttp.PluginQueue)
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody(ihttp.AddReg[0]+"插件已被激活", order.P.FromWxId))
				}
			}
		})
		return ihttp.FuncFinish(order.Name, order.Cron)
	})
	ihttp.AddPlugin(func(order ihttp.AddOrder) interface{} {
		order.Cron = ""
		order.Name = "查找插件"
		order.RegStr = "^命令$"
		order.Admin = true
		order.RegBool = true
		order.DailyFunction(func() {
			var msg string
			msg = "当前加载全部插件：\n"
			for _, i := range ihttp.PluginNameList {
				msg += i + "\n"
			}
			ihttp.PostIHttp(
				ihttp.BuildSendTextMsgBody(msg, order.P.FromWxId))
		})
		return ihttp.FuncFinish(order.Name, order.Cron)
	})
	ihttp.AddPlugin(func(order ihttp.AddOrder) interface{} {
		order.Cron = ""
		order.Name = "关闭插件"
		order.RegStr = "关闭(.*)"
		order.Admin = true
		order.RegBool = true
		order.DailyFunction(func() {
			for j, i := range ihttp.PluginNameList {
				if i == ihttp.AddReg[0] {
					ihttp.PluginQueue = ihttp.RemoveParam(ihttp.PluginQueue, ihttp.PluginList[j])
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody(ihttp.AddReg[0]+"插件已关闭", order.P.FromWxId))
					fmt.Println(j)
				}
			}
		})
		return ihttp.FuncFinish(order.Name, order.Cron)
	})
}
