package outside

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/silenceper/wechat/util"
	"school-helper/store"
	"time"
)

type DailyEnglish struct {
	Content     string `json:"content"`
	Note        string `json:"translation"` //Fuck，Note and Translation fields is not ordered.
	Translation string `json:"note"`
	Picture     string `json:"picture"`
	Dateline    string `json:"dateline"`
}

var keyPrefix = "outside:daily_english:"

func GetDailyEnglish() DailyEnglish {
	date := "" //TODO If date is empty, it means today

	tf := time.Now().Format("2006-01-02")
	redisKey := keyPrefix + tf

	result := new(DailyEnglish)
	//Retry 3 times
	for i := 0; i < 3; i++ {
		data := get(redisKey, date)

		err := json.Unmarshal(data, result)
		if err != nil {
			log.Fatal(err)
		}

		if validate(result) {
			return *result
		} else {
			clear(redisKey)
		}
	}

	return *result
}

func get(key, date string) []byte {
	data, err := store.RedisClient.Get(key).Result()

	var ret []byte

	if err == store.RedisNil || err != nil {
		//The key does not exists
		ret = getDailyEnglishFromAPI(date)
		set(key, ret)
		return ret
	} else {
		return []byte(data)
	}

}

//Get data from API
func getDailyEnglishFromAPI(date string) []byte {
	apiUrl := "http://open.iciba.com/dsapi?" + date
	ret, err := util.HTTPGet(apiUrl)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return ret
}

func validate(d *DailyEnglish) bool {
	if len(d.Content) > 0 && len(d.Dateline) > 0 && len(d.Translation) > 0 && len(d.Picture) > 0 {
		return true
	}
	return false

}

func set(key string, data []byte) (err error) {
	err = store.RedisClient.Set(key, string(data), 0).Err()
	return
}

func clear(key string) {
	store.RedisClient.Del(key).Err()
}
