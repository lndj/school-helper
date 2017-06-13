package wechat

import "github.com/silenceper/wechat/message"

func ReplyNews(title, description, pic_url, url string) *message.Reply {

	articles := make([]*message.Article, 1)

	article := new(message.Article)
	article.Title = title
	article.Description = description
	article.PicURL = pic_url
	article.URL = url

	articles[0] = article

	news := message.NewNews(articles)
	return &message.Reply{message.MsgTypeNews, news}

}
