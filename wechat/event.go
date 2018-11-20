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
	picUrl := "http://ww1.sinaimg.cn/large/65209136gw1f7vhjw95eqj20wt0zk40z.jpg"

	url := "https://github.com/lndj"
	return ReplyNews("小助手正在规划中", "请期待一个全新的小助手吧，如果你有好的想法，请留言给小助手", picUrl, url)

}

//The CiBa api, send by News
func eventDailyEnglish() *message.Reply {
	de := outside.GetDailyEnglish()
	desc := de.Content + "\n\n" + de.Translation
	title := "每日一句  " + de.Dateline

	return ReplyNews(title, desc, de.Picture, "")

}
