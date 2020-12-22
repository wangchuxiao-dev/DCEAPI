package exchange

import (
	"fmt"
	"net/http"
	_ "io/ioutil"
	"net/url"
)

type exchange struct {
	Name string
	Apikey string
	Secert string
	Password string
	LimitRate bool
}

func BaseRequest(method string, path string, params url.Values) {
	url, err := url.Parse(path)
	if err != nil {
		fmt.Println("invalid", url)
	}
	if method == "post" {
		res, err := http.Post(urlPath)
	} else {
		url.RawQuery = params.Encode()
		urlPath := url.String()
		res, err := http.Get(urlPath)
	}
}