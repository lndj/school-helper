package outside

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/silenceper/wechat/util"
)

type DailyEnglish struct {
	Content     string `json:"content"`
	Note        string `json:"translation"` //神坑，此处字段是反过来的，我也很无奈啊[无辜脸]
	Translation string `json:"note"`
	Picture     string `json:"picture"`
	Dateline    string `json:"dateline"`
}

func GetDailyEnglish() DailyEnglish {
	date := "" //TODO 此处为空则默认为当天
	data := getDailyEnglishFromAPI(date)

	var result DailyEnglish
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

//TODO 加入数据校验之后，每日只需要请求一次API，其余从redis直接获取

//从API获取数据
func getDailyEnglishFromAPI(date string) []byte {
	api_url := "http://open.iciba.com/dsapi?" + date
	ret, err := util.HTTPGet(api_url)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return ret
}
