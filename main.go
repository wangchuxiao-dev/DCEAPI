package main

import (
	"./exchange"
	"fmt"
)

func main() {
	res, _ := exchange.BaseRequest("GET", "https://www.baidu.com")
	fmt.Println(res)
}