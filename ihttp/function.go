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
	"wechat_http/util"
)

type PostBody[T any] struct {
	WxEvent         string `json:"event"`
	WxWx_RobotWxId  string `json:"robot_wxid"`
	WxRobotName     string `json:"robot_name"`
	WxType          int    `json:"type"`
	WxFromWxId      string `json:"from_wxid"`
	WxFromName      string `json:"from_name"`
	WxFinalFromWxId string `json:"final_from_wxid"`
	WxFinalFromName string `json:"final_from_name"`
	WxToWxId        string `json:"to_wxid"`
	WxMsgId         string `json:"msgid"`
	WxMsg           any    `json:"msg"`
	WxGroupWxId     string `json:"group_wxid"`

	Tgmsg_from_id  any `json:"Tg_msg_from_id"`
	Tgchanelname   any `json:"Tg_chanel_name"`
	TgMsg          any `json:"Tg_msg"`
	TgIfBot        any `json:"Tg_ifbot"`
	TgGroup        any `json:"Tg_ifgroup"`
	TgChanel       any `json:"Tg_ifchanel"`
	TgMsgfirstname any `json:"Tg_first_name"`
	TgmsgleastName any `json:"Tg_least_name"`
	TgUsername     any `json:"Tg_username"` //这里拿到tg传过来的信息，需要处理。
}

type AddOrder struct {
	Cron    string //这里是cron表达式，只有cron表达式才能被调用
	Method  func()
	RegStr  string //这里是正则匹配该信息为true才执行
	RegBool bool   //这里需要匹配信息为true才判断执行，false
	Name    string //这里是插件的名字
	Admin   bool   //这里管理员才执行，false为所有人都执行
	P       PostBody[any]
}

var Ch = make(chan PostBody[any], 1)
var Con bool
var AddReg []string
var OtMsg any
var OtmsgMasters string
var FromMsg any

func (bot AddOrder) DailyFunction(d func()) {
	if bot.P.WxMsg != nil {
		OtMsg = bot.P.WxMsg
		OtmsgMasters = "wxid_xq2w7jl6cbi811"
		FromMsg = bot.P.WxFinalFromWxId
	} else {
		OtMsg = bot.P.TgMsg
		OtmsgMasters = "1123322058"
		FromMsg = bot.P.Tgmsg_from_id
	}
	pMsg := util.StrVal(OtMsg) //判断哪个不为空则执行哪个。当微信信息为null，telegram不为空
	var msg string
	msg = "插件名字:" + bot.Name
	msg += "\nRegStr：" + bot.RegStr
	bot.Method = d

	if strings.Contains(fmt.Sprintf("%v", bot.Cron), " *") {
		bot.CronFunction(d) //去除这免得每次运行都判断一次。。
	}
	match, _ := regexp.MatchString(bot.RegStr, pMsg)
	if match != true || bot.RegBool != true {
		return
	}

	if bot.Admin == true {
		if FromMsg != OtmsgMasters {
			return
		}
	}
	re := regexp.MustCompile(bot.RegStr)
	matchArr := re.FindStringSubmatch(pMsg)

	for _, i := range matchArr[1:] {
		AddReg = append(AddReg, i)
	}
	go func() {
		defer func() {
			AddReg = []string{}
		}()
		bot.Method()
	}()

}

var Cronfunmap map[string]*cron.Cron

func init() {
	Cronfunmap = make(map[string]*cron.Cron)
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
	Cronfunmap[bot.Name] = c
	for _, cron_Func := range Cronfunmap {
		cron_Func.Start()
	}
}

func (bot AddOrder) Await(sender chan PostBody[any], timeout time.Duration, fromWxid string) interface{} {
	<-Ch
	Con = false
	PostIHttp(
		BuildSendTextMsgBody("给你"+timeout.String()+"秒的时间来回复", bot.P.WxFromWxId))
	for {
		select {
		case message := <-sender:
			if fromWxid == message.WxFromWxId {
				//if isDigit(message.Msg) {
				//	return message.Msg
				if message.WxMsg == "结束" {
					fmt.Println("结束")
					sender <- message
					return "结束"
				} else if message.WxMsg == "y" || message.WxMsg == "Y" {
					fmt.Println("yes")
					sender <- message
					return "yes"
				} else if message.WxMsg == "n" || message.WxMsg == "N" {
					fmt.Println("no")
					sender <- message
					return "no"
				} else {
					sender <- message
					//fmt.Println(message.Msg)
					PostIHttp(
						BuildSendTextMsgBody("输入错误！", bot.P.WxFromWxId))
				}
			} else {
				sender <- message
			}

		case <-time.After(time.Second * timeout):
			PostIHttp(
				BuildSendTextMsgBody(timeout.String()+"秒没有收到消息，正在为你结束对话\n",
					bot.P.WxFromWxId))
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
