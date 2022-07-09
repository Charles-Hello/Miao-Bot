package plugins

import (
	"wechat_http/ihttp"
	"wechat_http/util"
)

func init() {
	ihttp.AddPlugin(func(order ihttp.AddOrder) interface{} {
		order.Cron = ""
		order.Name = "同意朋友"
		order.RegStr = ""
		order.RegBool = order.P.Event == ihttp.EventFriendVerify
		order.Admin = true
		//这里我们要正则干嘛？
		order.DailyFunction(func() {
			// 同意好友请求
			pMsg := util.StrVal(order.P.Msg)
			ihttp.PostIHttp(ihttp.BuildAgreeFriendVerifyMsgBody(pMsg, order.P.ToWxId))
			// 发送消息
			ihttp.PostIHttp(
				ihttp.BuildSendTextMsgBody("Hello < "+order.P.FromName+" >, 稳我有事吗？!",
					order.P.FromWxId))
		})
		return ihttp.FuncFinish(order.Name, order.Cron)
	})
}
