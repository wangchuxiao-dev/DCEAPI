package main

import (
	"github.com/PythonohtyP1900/DCEAPI"
	"github.com/PythonohtyP1900/DCEAPI/exchanges"
	"fmt"
)

func main() {
	aofex := exchanges.NewAofex(false, "testsecret", "testapikey")
	s, err := aofex.FetchMarkets()
	fmt.Println(err, s)
	_, err2 := DCEAPI.BaseRequest("POST", "https://www.baidu.com", "")
	fmt.Println(err2)
}