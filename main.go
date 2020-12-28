package main

import (
	"./exchange"
	"fmt"
)

func main() {
	aofex := &exchange.Aofex{}
	res, err := exchange.BaseRequest("POST", "https://www.baidu.cm", "")
	fmt.Println(res, err)
}