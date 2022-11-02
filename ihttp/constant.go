package ihttp

// msgBody 事件类型常量
const (
	// EventSendOutMsg 机器人主动消息事件
	EventSendOutMsg = "EventSendOutMsg"
	// EventLogin 机器人上线
	EventLogin = "EventLogin"
	// EventGroupMsg 群消息事件
	EventGroupMsg = "EventGroupMsg"
	// EventFriendMsg 私聊消息事件
	EventFriendMsg = "EventFriendMsg"
	// EventFriendVerify 新的好友申请
	EventFriendVerify = "EventFriendVerify"
	// AgreeFriendVerify 同意好友申请
	AgreeFriendVerify = "AgreeFriendVerify"
	// SendTextMsg 发送文本消息 Wx_RobotWxId toWxId(群/好友) msg
	SendTextMsg = "SendTextMsg"
	// SendImageMsg 发送图片消息 Wx_RobotWxId toWxId(群/好友) msg(name[md5值或其他唯一的名字，包含扩展名例如1.jpg], url)
	SendImageMsg = "SendImageMsg"
	// SendLinkMsg 发送分享链接 Wx_RobotWxId, toWxId(群/好友), msg(title, text, target_url, pic_url, icon_url)
	SendLinkMsg = "SendLinkMsg"
	// SendCardMsg 发送名片 Wx_RobotWxId, toWxId(群/好友), msg(你要发的那个名片的 wxid)
	SendCardMsg = "SendCardMsg"
	// RevokeMsg 撤回消息API (Wx_RobotWxId, msg)
	RevokeMsg = "RevokeMsg"
	// GetWechatMoments 获取朋友圈(Wx_RobotWxId, msg)
	GetWechatMoments = "GetWechatMoments"
	// ReplyWechatMoments 评论朋友圈(Wx_RobotWxId, msg[moments_id,content])
	ReplyWechatMoments = "ReplyWechatMoments"
	// SendDiyMusicMsg 发送自定义音乐 (Wx_RobotWxId, toWxId,msg[name, singer, home, url, type])
	SendDiyMusicMsg = "SendDiyMusicMsg"
)

// 消息类型常量
const (
	// MsgTypeText 文本消息
	MsgTypeText = 1
	// MsgTypeImg 图片消息
	MsgTypeImg = 3
	// MsgTypeVoice 语音消息
	MsgTypeVoice = 34
	// MsgTypeNewFriend 好友申请
	MsgTypeNewFriend = 37
	// MsgTypeCarte 名片消息
	MsgTypeCarte = 42
	// MsgTypeRevoke 撤回消息
	MsgTypeRevoke = 2005

	//43/视频 47/动态表情 48/地理位置  49/分享链接
	//2000/转账 2001/红包  2002/小程序  2003/群邀请
)
