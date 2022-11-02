package plugins

import (
	"wechat_http/ihttp"
	"wechat_http/util"
)

func init() {
	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "同意朋友"
		configs.RegStr = ""
		configs.RegBool = configs.P.WxEvent == ihttp.EventFriendVerify
		configs.Admin = true
		//这里我们要正则干嘛？
		configs.DailyFunction(func() {
			// 同意好友请求
			pMsg := util.StrVal(configs.P.WxMsg)
			ihttp.PostIHttp(ihttp.BuildAgreeFriendVerifyMsgBody(pMsg, configs.P.WxToWxId))
			// 发送消息
			ihttp.PostIHttp(
				ihttp.BuildSendTextMsgBody("Hello < "+configs.P.WxFromName+" >, 稳我有事吗？!",
					configs.P.WxFromWxId))
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})
}
