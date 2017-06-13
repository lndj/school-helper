package outside

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ReplyContent struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func ReplyByRobot(content, uid string) ReplyContent {
	ret := getReplyFromTuling(content, uid)
	var result ReplyContent
	err := json.Unmarshal(ret, &result)
	if err != nil {
		return ReplyContent{
			100000,
			"啊，我不知道说啥了",
		}
	}

	return result
}

func getReplyFromTuling(content, uid string) []byte {
	api_url := "http://www.tuling123.com/openapi/api"
	resp, err := http.PostForm(api_url, url.Values{
		"key":    {"6019ba8b60b62ad10d695aed384010df"},
		"info":   {content},
		"userid": {uid},
	})
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return body

}
