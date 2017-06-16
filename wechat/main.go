package wechat

import (
	"fmt"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"gopkg.in/gin-gonic/gin.v1"
	"school-helper/config"
)

//处理微信相关的逻辑
//wechat 模块下分类型处理各项事件
//之后微信接受到消息，需要处理业务逻辑，分相应的业务到相应的模块
func WechatHandler(c *gin.Context) {
	//配置微信参数
	appID, _ := config.Configure.String("WECHAT_APP_ID")
	appSecret, _ := config.Configure.String("WECHAT_APP_SECRET")
	token, _ := config.Configure.String("WECHAT_TOKEN")
	encodingAESKey, _ := config.Configure.String("WECHAT_AES_KEY")

	wechatOption := &wechat.Config{
		AppID:          appID,
		AppSecret:      appSecret,
		Token:          token,
		EncodingAESKey: encodingAESKey,
	}
	wc := wechat.NewWechat(wechatOption)

	// 传入request和responseWriter
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

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

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
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
