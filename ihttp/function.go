package ihttp

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
	"wechat_http/config"
)

type PostBody struct {
	Event         string `json:"event"`
	RobotWxId     string `json:"robot_wxid"`
	RobotName     string `json:"robot_name"`
	Type          int    `json:"type"`
	FromWxId      string `json:"from_wxid"`
	FromName      string `json:"from_name"`
	FinalFromWxId string `json:"final_from_wxid"`
	FinalFromName string `json:"final_from_name"`
	ToWxId        string `json:"to_wxid"`
	MsgId         string `json:"msgid"`
	Msg           string `json:"msg"`
	GroupWxId     string `json:"group_wxid"`
}

type AddOrder struct {
	Cron    string //这里是cron表达式，只有cron表达式才能被调用
	Method  func()
	RegStr  string //这里是正则匹配该信息为true才执行
	RegBool bool   //这里需要匹配信息为true才判断执行，false
	Name    string //这里是插件的名字
	Admin   bool   //这里管理员才执行，false为所有人都执行
	P       PostBody
}

var Ch = make(chan PostBody, 1)
var Con bool
var AddReg []string

func (bot AddOrder) DailyFunction(d func()) {
	var msg string
	msg = "插件名字:" + bot.Name
	msg += "\nRegStr：" + bot.RegStr
	bot.Method = d
	if strings.Contains(fmt.Sprintf("%v", bot.Cron), " *") {
		bot.CronFunction(d) //去除这免得每次运行都判断一次。。
	}
	match, _ := regexp.MatchString(bot.RegStr, bot.P.Msg)
	if match != true || bot.RegBool != true {
		return
	}
	if bot.Admin == true {
		if bot.P.FinalFromWxId != config.Config.IHttp.MasterWxId {
			return
		}
	}
	re := regexp.MustCompile(bot.RegStr)
	matchArr := re.FindStringSubmatch(bot.P.Msg)

	for _, i := range matchArr[1:] {
		AddReg = append(AddReg, i)
	}
	go bot.Method()
	fmt.Println(msg)
	AddReg = []string{}
}

func (bot AddOrder) CronFunction(d func()) {
	var msg string
	msg = "定时任务执行中！\n"
	msg += "cron名字：" + bot.Name + "\n"
	msg += "cron时间:" + bot.Cron
	fmt.Println(msg)
	c := cron.New()
	_, err := c.AddFunc(bot.Cron, d)
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}

func (bot AddOrder) Await(sender chan PostBody, timeout time.Duration, fromWxid string) interface{} {
	<-Ch
	Con = false
	PostIHttp(
		BuildSendTextMsgBody("给你"+timeout.String()+"秒的时间来回复", bot.P.FromWxId))
	for {
		select {
		case message := <-sender:
			if fromWxid == message.FromWxId {
				//if isDigit(message.Msg) {
				//	return message.Msg
				if message.Msg == "结束" {
					fmt.Println("结束")
					sender <- message
					return "结束"
				} else if message.Msg == "y" || message.Msg == "Y" {
					fmt.Println("yes")
					sender <- message
					return "yes"
				} else if message.Msg == "n" || message.Msg == "N" {
					fmt.Println("no")
					sender <- message
					return "no"
				} else {
					sender <- message
					//fmt.Println(message.Msg)
					PostIHttp(
						BuildSendTextMsgBody("输入错误！", bot.P.FromWxId))
				}
			} else {
				sender <- message
			}

		case <-time.After(time.Second * timeout):
			PostIHttp(
				BuildSendTextMsgBody(timeout.String()+"秒没有收到消息，正在为你结束对话\n",
					bot.P.FromWxId))
			return nil
		}
	}

}

type PluginFunc func(AddOrder) interface{}

var PluginNameList []string
var PluginList []PluginFunc

var PluginQueue []PluginFunc

// AddPlugin 将插件放入队列
func AddPlugin(p PluginFunc) {
	RePlugin := p(AddOrder{})
	if strings.Contains(fmt.Sprintf("%v", RePlugin), " *") { //这里如果cron有值，则排除。cron没有值则拿到name进行命令
	} else {
		PluginList = append(PluginList, p)
		PluginNameList = append(PluginNameList, RePlugin.(string))
		PluginQueue = append(PluginQueue, p)
	}
}

func RemoveParam(sli []PluginFunc, n PluginFunc) []PluginFunc {
	for i := 0; i < len(sli); i++ {
		if GetFunctionName(sli[i]) == GetFunctionName(n) {
			if i == 0 {
				sli = sli[1:]
			} else if i == len(sli)-1 {
				sli = sli[:i]
			} else {
				sli = append(sli[:i], sli[i+1:]...)
			}
			i-- // 如果索引i被去掉后，原来索引i+1的会移动到索引i
		}
	}
	return sli
}

func GetFunctionName(i interface{}, seps ...rune) string {
	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

func FuncFinish(name, cron string) string {
	if strings.Contains(fmt.Sprintf("%v", cron), " *") {
		return cron
	} else {
		return name
	}
}
