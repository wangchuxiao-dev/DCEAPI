package exchanges

import (
	"github.com/PythonohtyP1900/DCEAPI"

	"time"
)

func NewExchange(exchange, apiKey, secret string) DCEAPI.SpotAPI {
	var api DCEAPI.SpotAPI
	switch exchange {
	case "aofex":
		api = NewAofex(secret, apiKey)
	case "huobipro":
		api = NewHuobi(secret, apiKey)
	default:
		panic("Unsupport Exchange")
	}
	return api
}

// 时间转为时间戳
func dateTimeToTimeStamp(datetime string) int {
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
	timeUnix := tmp.Unix()
	return int(timeUnix)
}
