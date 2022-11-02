package plugins

import (
	"fmt"
	"wechat_http/ihttp"
)

func init() {
	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "激活插件"
		configs.RegStr = "^激活(.*)"
		configs.Admin = true
		configs.RegBool = true
		configs.DailyFunction(func() {
			for j, i := range ihttp.PluginNameList {
				if i == ihttp.AddReg[0] {
					ihttp.PluginQueue = append(ihttp.PluginQueue, ihttp.PluginList[j])
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody(ihttp.AddReg[0]+"插件已被激活", configs.P.WxFromWxId))
				}
			}
			for cronName, cron_Func := range ihttp.Cronfunmap {
				if ihttp.AddReg[0] == cronName {
					cron_Func.Start()
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody(cronName+"插件已被激活", configs.P.WxFromWxId))
				}
			}
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})
	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "查找插件"
		configs.RegStr = "^命令$"
		configs.Admin = true
		configs.RegBool = true
		configs.DailyFunction(func() {
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
				ihttp.BuildSendTextMsgBody(msg, configs.P.WxFromWxId))
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})
	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "关闭插件"
		configs.RegStr = "关闭(.*)"
		configs.Admin = true
		configs.RegBool = true
		configs.DailyFunction(func() {
			for j, i := range ihttp.PluginNameList {
				if i == ihttp.AddReg[0] {
					ihttp.PluginQueue = ihttp.RemoveParam(ihttp.PluginQueue, ihttp.PluginList[j])
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody(ihttp.AddReg[0]+"插件已关闭", configs.P.WxFromWxId))
				}
			}
			for cronName, cronFunc := range ihttp.Cronfunmap {
				if ihttp.AddReg[0] == cronName {
					cronFunc.Stop()
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody(cronName+"插件已关闭", configs.P.WxFromWxId))
				}
			}
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})
}
