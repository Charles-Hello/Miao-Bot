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
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody(ihttp.AddReg[0]+"插件已被激活", order.P.FromWxId))
				}
			}
			for cronName, cron_Func := range ihttp.Cronfunmap {
				if ihttp.AddReg[0] == cronName {
					cron_Func.Start()
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody(cronName+"插件已被激活", order.P.FromWxId))
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
			for cronName := range ihttp.Cronfunmap {
				fmt.Println(cronName)
				msg += "『定时』" + cronName + "\n"
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
				}
			}
			for cronName, cronFunc := range ihttp.Cronfunmap {
				if ihttp.AddReg[0] == cronName {
					cronFunc.Stop()
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody(cronName+"插件已关闭", order.P.FromWxId))
				}
			}
		})
		return ihttp.FuncFinish(order.Name, order.Cron)
	})
}
