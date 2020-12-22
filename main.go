package main

import (
	_ "fmt"
	"./exchange"
	"net/url"
)

func main() {
	params = url.Values{}
	exchange.BaseRequest("GET", "www.huobipro.com", nil)
}