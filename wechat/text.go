package wechat

import (
	"github.com/silenceper/wechat/message"

	"github.com/lndj/school-helper/outside"
)

//The text message router
func ReplyText(content string) *message.Reply {
	text := message.NewText(content)
	return &message.Reply{message.MsgTypeText, text}
}

//Replied by Tuling api
func ReplyByRobot(content, openid string) *message.Reply {
	//获取机器人返回的回复
	ret := outside.ReplyByRobot(content, openid)
	text := message.NewText(ret.Text)
	return &message.Reply{message.MsgTypeText, text}
}
