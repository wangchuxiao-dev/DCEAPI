package exchanges

import (
	"github.com/PythonohtyP1900/DCEAPI"
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