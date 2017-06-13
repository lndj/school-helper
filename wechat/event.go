package wechat

import (
	"github.com/silenceper/wechat/message"
	"study/outside"
)

func Subscribe() *message.Reply {
	content := "欢迎你的关注哦！"
	return ReplyText(content)

}

func Click(v message.MixMessage) *message.Reply {
	if v.EventKey == "每日一句" {
		return eventDailyEnglish()
	}

	pic_url := "http://ww1.sinaimg.cn/large/65209136gw1f7vhjw95eqj20wt0zk40z.jpg"
	url := "https://github.com/lndj"
	return ReplyNews("我的GitHub", "这就是我的GayHub哦", pic_url, url)

}

//金山词霸每日一句
//图文方式发送
func eventDailyEnglish() *message.Reply {
	de := outside.GetDailyEnglish()
	desc := de.Content + "\n\n" + de.Translation
	title := "每日一句  " + de.Dateline

	return ReplyNews(title, desc, de.Picture, "")

}
