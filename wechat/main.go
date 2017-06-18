package wechat

import (
	"fmt"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/lndj/school-helper/config"
)

//Wechat handler, handle all of the task about wechat
func WechatHandler(c *gin.Context) {
	wechatOption := &wechat.Config{
		AppID:          config.Environment.WechatAppID,
		AppSecret:      config.Environment.WechatAppSecret,
		Token:          config.Environment.WechatToken,
		EncodingAESKey: config.Environment.WechatAesKey,
	}
	wc := wechat.NewWechat(wechatOption)

	// Param is request and responseWriter
	server := wc.GetServer(c.Request, c.Writer)
	//set message handler
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//Get user's openid
		openid := server.GetOpenID()

		switch msg.MsgType {
		case message.MsgTypeText:
			return ReplyByRobot(msg.Content, openid)
		case message.MsgTypeEvent:
			return eventHandler(msg)
		case message.MsgTypeVoice:
			return ReplyText("你的声音真难听！")
		}

		return nil
	})

	//Handle message receive and reply
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Send the message
	server.Send()
}

func eventHandler(v message.MixMessage) *message.Reply {

	switch v.Event {
	case message.EventSubscribe:
		return Subscribe()
	case message.EventScan:
		return nil
	case message.EventUnsubscribe:
		return nil
	case message.EventLocation:
		return nil
	case message.EventClick:
		return Click(v)
	case message.EventView:
		return Click(v)
	case message.EventScancodePush:
		return nil

		// 扫码推事件且弹出“消息接收中”提示框的事件推送
	case message.EventScancodeWaitmsg:
		return nil

		// 弹出系统拍照发图的事件推送
	case message.EventPicSysphoto:
		return nil

		// 弹出拍照或者相册发图的事件推送
	case message.EventPicPhotoOrAlbum:
		return nil

		// 弹出微信相册发图器的事件推送
	case message.EventPicWeixin:
		return nil

		// 弹出地理位置选择器的事件推送
	case message.EventLocationSelect:
		return nil
	default:
		return nil
	}

}
