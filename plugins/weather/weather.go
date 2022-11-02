package weather

import (
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"unicode/utf8"
	"wechat_http/ihttp"
	"wechat_http/util"
)

func TQPlugin[T any](msgBody *ihttp.PostBody[any]) {
	val := util.StrVal(msgBody.WxMsg)
	str := []rune(val)
	l := utf8.RuneCountInString(val)
	if l < 4 || l > 6 {
		return
	}
	// 天气结尾
	endStr := string(str[l-2:])

	if endStr == "天气" {
		// len(字符串变量)获取字符串的字节长度,其中英文占1个字节长度,中文占用3个字节长度
		city := string(str[:l-2])
		weather := GetWeather(city)
		var body ihttp.PostBody[any]
		if msgBody.WxEvent == ihttp.EventGroupMsg {
			// 艾特回复
			body = ihttp.BuildSendGroupMsgAndAt("\n"+weather, msgBody.WxFromWxId,
				msgBody.WxFinalFromWxId, msgBody.WxFinalFromName)

		} else if msgBody.WxEvent == ihttp.EventFriendMsg {
			body = ihttp.BuildSendTextMsgBody[any](weather, msgBody.WxFromWxId)
		}
		ihttp.PostIHttp(body)
	}
}

type Err struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Success struct {
	CityId     string `json:"cityid"`
	City       string `json:"city"`
	CityEn     string `json:"cityEn"`
	Country    string `json:"country"`
	CountryEn  string `json:"countryEn"`
	UpdateTime string `json:"update_time"`
	Data       []struct {
		Day           string        `json:"day"`
		Date          string        `json:"date"`
		Week          string        `json:"week"`
		Wea           string        `json:"wea"`
		WeaImg        string        `json:"wea_img"`
		WeaDay        string        `json:"wea_day"`
		WeaDayImg     string        `json:"wea_day_img"`
		WeaNight      string        `json:"wea_night"`
		WeaNightImg   string        `json:"wea_night_img"`
		Tem           string        `json:"tem"`
		Tem1          string        `json:"tem1"`
		Tem2          string        `json:"tem2"`
		Humidity      string        `json:"humidity"`
		Visibility    string        `json:"visibility"`
		Pressure      string        `json:"pressure"`
		Win           []string      `json:"win"`
		WinSpeed      string        `json:"win_speed"`
		WinMeter      string        `json:"win_meter"`
		Sunrise       string        `json:"sunrise"`
		Sunset        string        `json:"sunset"`
		Air           string        `json:"air"`
		AirLevel      string        `json:"air_level"`
		AirTips       string        `json:"air_tips"`
		Phrase        string        `json:"phrase"`
		Narrative     string        `json:"narrative"`
		Moonrise      string        `json:"moonrise"`
		MoonSet       string        `json:"moonset"`
		MoonPhrase    string        `json:"moonPhrase"`
		Rain          string        `json:"rain"`
		UvIndex       string        `json:"uvIndex"`
		UvDescription string        `json:"uvDescription"`
		Alarm         []interface{} `json:"alarm,omitempty"`
	} `json:"data"`
	Nums int `json:"nums"`
}

func GetWeather(cityName string) string {
	url := fmt.Sprintf("https://v0.yiketianqi.com/api?unescape=1&version=v91&appid=43656176&appsecret=I42og6Lm&ext=&cityid=&city=%s", cityName)
	client := &http.Client{}
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "查询天气失败!"
	}
	response, err := client.Do(r)
	body, err := ioutil.ReadAll(response.Body)
	var str Err
	jsonIter.Unmarshal(body, &str)
	fmt.Println(str)
	if str.ErrCode == 100 {
		return fmt.Sprintf("%s > 不存在,天气查询失败", cityName)
	}
	var res Success
	jsonIter.Unmarshal(body, &res)
	fmt.Println(res)
	a := fmt.Sprintf("小爱天气助手 > %s \n城市: %s \n更新时间:  %s \n天气: %s \n当前体感温度: %s°C \n最高气温: %s°C  \n最低气温: %s°C ",
		res.Data[0].Week, res.City, res.UpdateTime, res.Data[0].Wea, res.Data[0].Tem,
		res.Data[0].Tem1, res.Data[0].Tem2)
	return a
}
