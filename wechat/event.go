package wechat

import (
	"github.com/silenceper/wechat/message"

	"github.com/lndj/school-helper/outside"
)

//When user subscribe, run this
func Subscribe() *message.Reply {
	content := "欢迎你的关注哦！"
	return ReplyText(content)

}

//Click event handler
func Click(v message.MixMessage) *message.Reply {
	if v.EventKey == "每日一句" {
		return eventDailyEnglish()
	}

	pic_url := "http://ww1.sinaimg.cn/large/65209136gw1f7vhjw95eqj20wt0zk40z.jpg"
	url := "https://github.com/lndj"
	return ReplyNews("我的GitHub", "这就是我的GayHub哦", pic_url, url)

}

//The CiBa api, send by News
func eventDailyEnglish() *message.Reply {
	de := outside.GetDailyEnglish()
	desc := de.Content + "\n\n" + de.Translation
	title := "每日一句  " + de.Dateline

	return ReplyNews(title, desc, de.Picture, "")

}
