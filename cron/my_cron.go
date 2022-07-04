package cron

//
//import (
//	"github.com/robfig/cron/v3"
//	"log"
//	"wechat_http/ihttp/service"
//	"wechat_http/plugins/weather"
//)
//
//var WorkRemind *cron.Cron
//
//func InitMyCron() {
//	c := cron.New()
//
//	// 定时天气提醒
//	spec := "00 07 * * *"
//	addFunc, err := c.AddFunc(spec, func() {
//		getWeather := weather.GetWeather("郑州")
//		service.SendMsgToMe(getWeather)
//	})
//	if err != nil {
//		log.Println(addFunc, err)
//	}
//	c.Start()
//
//}
//
//// goToWorkRemind 上班打卡提醒
//func goToWorkRemind() {
//	c := cron.New()
//	//spec := "@every 3s"
//	spec1 := "25 08 * * *"
//	c.AddFunc(spec1, func() {
//		service.SendMsgToMe("上班了 不要忘记钉打卡!")
//	})
//
//	spec2 := "35 17 * * *"
//	c.AddFunc(spec2, func() {
//		service.SendMsgToMe("下班了 不要忘记钉打卡!")
//	})
//	WorkRemind = c
//}
//
//func init() {
//	//goToWorkRemind()
//}
