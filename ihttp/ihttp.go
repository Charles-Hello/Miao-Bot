package ihttp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"wechat_http/config"
)

// PostBody 基础post消息体(收发消息通用) (与iHttp插件进行交互)

// LinkMsgBody 链接消息体
type LinkMsgBody struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	TargetUrl string `json:"target_url"`
	PicUrl    string `json:"pic_url"`
	IconUrl   string `json:"icon_url"`
}

func PostIHttp(body PostBody[any]) {

	body.RobotWxId = config.Config.IHttp.RobotWxId
	bytesData, _ := json.Marshal(body)
	client := &http.Client{}
	//提交请求
	req, err := http.NewRequest("POST", config.Config.IHttp.URL, bytes.NewReader(bytesData))
	//增加header选项
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", config.Config.IHttp.Authorization)
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(req)
	defer response.Body.Close()
}

// BuildSendTextMsgBody 构建发送普通文本消息
func BuildSendTextMsgBody(msg, toWxId string) PostBody[any] {
	body := PostBody[any]{
		Event:  SendTextMsg,
		ToWxId: toWxId,
		Msg:    msg,
	}
	return body
}

// BuildSendGroupMsgAndAt 构建群聊消息并艾特
func BuildSendGroupMsgAndAt(msg, fromWxId, finalFromWxId, finalFromName string) PostBody[any] {
	body := PostBody[any]{
		Event:         "SendGroupMsgAndAt",
		ToWxId:        fromWxId,
		Msg:           msg,
		GroupWxId:     fromWxId,
		FinalFromName: finalFromName,
		FinalFromWxId: finalFromWxId,
	}
	return body
}

// BuildAgreeFriendVerifyMsgBody 构建同意好友申请
func BuildAgreeFriendVerifyMsgBody(msg, toWxId string) PostBody[any] {
	body := PostBody[any]{
		Event:  AgreeFriendVerify,
		Msg:    msg,
		ToWxId: toWxId,
	}
	return body
}

// BuildSendLinkMsg 构建 发送 卡片链接消息
func BuildSendLinkMsg(title, text, targetUrl, picUrl, IconUrl, toWxId string) PostBody[any] {
	body := PostBody[any]{
		Event:  "SendLinkMsg",
		ToWxId: toWxId,
		Msg: LinkMsgBody{
			Title:     title,
			Text:      text,
			TargetUrl: targetUrl,
			PicUrl:    picUrl,
			IconUrl:   IconUrl,
		},
	}
	return body
}

// BuildSendCardMsg 发送名片
func BuildSendCardMsg(cardWxId, toWxId string) PostBody[any] {
	body := PostBody[any]{
		Event:  "SendCardMsg",
		ToWxId: toWxId,
		Msg:    cardWxId,
	}
	return body
}

// BuildRevokeMsg 撤回消息
func BuildRevokeMsg(msgId string) PostBody[any] {
	body := PostBody[any]{
		Event: RevokeMsg,
		Msg:   msgId,
	}
	return body
}

// BuildInviteInGroup 构建邀请加群 groupId 群组id toWxId 被邀请人id
func BuildInviteInGroup(groupId, toWxId string) PostBody[any] {
	body := PostBody[any]{
		Event:     "InviteInGroup",
		ToWxId:    toWxId,
		GroupWxId: groupId,
	}
	return body
}

// BuildGetWechatMoments 获取朋友圈
func BuildGetWechatMoments(msg string) PostBody[any] {
	body := PostBody[any]{
		Event: "GetWechatMoments",
		Msg:   msg,
	}
	return body
}
