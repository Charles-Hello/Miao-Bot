package plugins

import (
	"strconv"
	"wechat_http/ihttp"
	"wechat_http/telegram_bot"
)

func init() {
	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "周杰伦你好"
		configs.RegStr = "周(.*)伦 ,你(.*)"
		configs.Admin = true
		configs.RegBool = true
		configs.DailyFunction(func() {
			for d, i := range ihttp.AddReg {
				ihttp.PostIHttp(
					ihttp.BuildSendTextMsgBody("第"+strconv.Itoa(d)+"个参数为："+i, configs.P.WxFromWxId))
				telegram_bot.Return.OpSlice = append(telegram_bot.Return.OpSlice, "你好")
			}
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})

	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "鸡你太美"
		configs.RegStr = "鸡你太美"
		configs.Admin = false
		configs.RegBool = true
		configs.DailyFunction(func() {
			telegram_bot.TG_send("好兄弟")
			telegram_bot.TG_sleep(2)
			telegram_bot.TG_edit("你怎么了")
			telegram_bot.TG_sleep(2)
			telegram_bot.TG_reply("芜湖")
			telegram_bot.TG_sleep(2)
			telegram_bot.TG_delete()
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})

	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "我是sad任务"
		configs.RegStr = "sad"
		configs.RegBool = true
		configs.Admin = true
		configs.DailyFunction(func() {
			ihttp.PostIHttp(
				ihttp.BuildSendTextMsgBody("我负责sad", configs.P.WxFromWxId))
			telegram_bot.Return.OpSlice = append(telegram_bot.Return.OpSlice, "我负责sad")
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})

	//ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
	//	configs.Cron = ""
	//	configs.Name = "我是撤回任务"
	//	configs.RegStr = "撤回"
	//	configs.RegBool = true
	//	configs.Admin = true
	//	configs.DailyFunction(func() {
	//		ihttp.PostIHttp(
	//			ihttp.BuildRevokeMsg(configs.P.MsgId))
	//	})
	//	return ihttp.FuncFinish(configs.Name, configs.Cron)
	//})

	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "我是2"
		configs.RegStr = "2"
		configs.RegBool = true
		configs.Admin = true
		configs.DailyFunction(func() {
			ihttp.PostIHttp(
				ihttp.BuildSendTextMsgBody("我是1", configs.P.WxFromWxId))
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})

	//ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
	//	configs.Cron = "*/12 * * * *"
	//	configs.Name = "我是定时任务"
	//	configs.DailyFunction(func() {
	//		ihttp.PostIHttp(
	//			ihttp.BuildSendTextMsgBody("我是定时任务", config.Config.Wx_MasterWxId))
	//	})
	//	return ihttp.FuncFinish(configs.Name, configs.Cron)
	//})
	//
	//ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
	//	configs.Cron = "*/10 * * * *"
	//	configs.Name = "我是第二个定时任务"
	//	configs.DailyFunction(func() {
	//		ihttp.PostIHttp(
	//			ihttp.BuildSendTextMsgBody("我是第二个定时任务", config.Config.Wx_MasterWxId))
	//	})
	//	return ihttp.FuncFinish(configs.Name, configs.Cron)
	//})
	//
	//ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
	//	configs.Cron = "*/3 * * * *"
	//	configs.Name = "我是第三个定时任务"
	//	configs.DailyFunction(func() {
	//		ihttp.PostIHttp(
	//			ihttp.BuildSendTextMsgBody("我是第三个定时任务", config.Config.Wx_MasterWxId))
	//	})
	//	return ihttp.FuncFinish(configs.Name, configs.Cron)
	//})

	ihttp.AddPlugin(func(configs ihttp.AddOrder) interface{} {
		configs.Cron = ""
		configs.Name = "指令"
		configs.RegStr = "测试指令"
		configs.RegBool = true
		configs.Admin = false
		configs.DailyFunction(func() {
			switch configs.Await(ihttp.Ch, 20, configs.P.WxFromWxId) {
			case "no":
				ihttp.PostIHttp(
					ihttp.BuildSendTextMsgBody("检测到你选择的是n，请输入「结束」",
						configs.P.WxFromWxId))
				switch configs.Await(ihttp.Ch, 20, configs.P.WxFromWxId) {
				case "结束":
					ihttp.PostIHttp(
						ihttp.BuildSendTextMsgBody("检测到你选择的是结束\n恭喜你，通关了",
							configs.P.WxFromWxId))
				}
			}
		})
		return ihttp.FuncFinish(configs.Name, configs.Cron)
	})
}
