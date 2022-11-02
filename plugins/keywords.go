package plugins

import (
	"wechat_http/ihttp"
)

func init() {
	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = "" // 如果不需要定时任务(DailyFunction)，则不需要设置,在DailyFunction中设置
		configs.Name = "ping"
		configs.RegStr = "^ping$"
		configs.RegBool = configs.P.WxEvent == ihttp.EventFriendMsg
		configs.Admin = true
		configs.DailyFunction(func() {
			ihttp.PostIHttp(
				ihttp.BuildSendTextMsgBody("pong[皱眉]", configs.P.WxFromWxId))
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})

	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "id"
		configs.RegStr = "^id$"
		configs.Admin = true
		configs.RegBool = true
		configs.DailyFunction(func() {
			ihttp.PostIHttp(
				ihttp.BuildSendTextMsgBody(configs.P.WxFromWxId, configs.P.WxFromWxId))
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})

	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "系统状态"
		configs.RegStr = "系统状态"
		configs.Admin = true
		configs.RegBool = true
		configs.DailyFunction(func() {
			ihttp.PostIHttp(ihttp.BuildSendTextMsgBody("还没写完", configs.P.WxFromWxId))
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})

	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "部署github"
		configs.RegStr = "部署地址"
		configs.Admin = true
		configs.RegBool = true
		configs.DailyFunction(func() {
			ihttp.PostIHttp(ihttp.BuildSendLinkMsg("github.com", "我的网站", "https://github.com/Charles-Hello/Miao-Bot",
				"https://cdn.wxy97.com/public/avatar.jpg", "https://cdn.wxy97.com/public/avatar.jpg", configs.P.WxFromWxId))
		})

		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})

	//if msg == "主人" {
	//	ihttp.PostIHttp(
	//		ihttp.BuildSendCardMsg(config.Config.IHttp.Wx_MasterWxId, msgBody.FromWxId))
	//}
	//if msg == "公众号" {
	//	ihttp.PostIHttp(
	//		ihttp.BuildSendCardMsg("gh_11e200c41d89", msgBody.FromWxId))
	//}
	//
	//if msgBody.FromWxId == config.Config.IHttp.Wx_MasterWxId {
	//	// 主人指令
	//	masterCommand(false, msg, msgBody.FromWxId, msgBody.FinalFromWxId, msgBody.FinalFromName)
	//}
	//publicCommand(false, msg, msgBody.FromWxId, msgBody.FinalFromWxId, msgBody.FinalFromName)

}

// 私人指令(群聊私聊通用)
//func masterCommand[T string](isGroup bool, msg, formWxId, finalFromWxId, finalFromName string) {
//	if msg == "指令" {
//		redis.Rdb.Set(redis.RdbCtx, "masterCommand:"+finalFromWxId, "", 120*time.Second)
//		command := "1.开启定时上班打卡提醒\n2.关闭定时上班打卡提醒\n[奸笑]请回复对应数字"
//		if isGroup {
//			ihttp.PostIHttp(ihttp.BuildSendGroupMsgAndAt("\n您的私人指令 >\n"+command, formWxId,
//				finalFromWxId, finalFromName))
//		} else {
//			ihttp.PostIHttp(ihttp.BuildSendTextMsgBody(command, formWxId))
//		}
//	}
//if msg == "我的博客" {
//	ihttp.PostIHttp(ihttp.BuildSendLinkMsg("王旭阳个人博客", "王旭阳个人博客", "https://www.wxy97.com/",
//		"https://cdn.wxy97.com/public/avatar.jpg", "https://cdn.wxy97.com/public/avatar.jpg", formWxId))
//}

//}

//// 公共指令(群聊私聊通用)
//func publicCommand(isGroup bool, msg, formWxId, finalFromWxId, finalFromName string) {
//	if msg == "运行环境" {
//		if isGroup {
//			ihttp.PostIHttp(ihttp.BuildSendGroupMsgAndAt("\n"+util.GetServerInfo(), formWxId,
//				finalFromWxId, finalFromName))
//		} else {
//			ihttp.PostIHttp(ihttp.BuildSendTextMsgBody(util.GetServerInfo(), formWxId))
//		}
//	}
//
//	if msg == "id" {
//		if isGroup {
//			ihttp.PostIHttp(ihttp.BuildSendGroupMsgAndAt("\n"+util.GetServerInfo(), formWxId,
//				finalFromWxId, finalFromName))
//		} else {
//			ihttp.PostIHttp(ihttp.BuildSendTextMsgBody(util.GetServerInfo(), formWxId))
//		}
//	}
//}
