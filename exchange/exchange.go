package exchange

import (
	"net/http"
	"strings"
	"io/ioutil"
	_ "net/url"
)

type Exchange struct {
	Name string
	instance interface{}
	Apikey string
	Secert string
	Password string
	LimitRate bool
}

func BaseRequest(method string, path string) (interface{}, error) {
	client := &http.Client{}
	if method == "GET" {
		fmt.Println("GET")
	} else {
		fmt.Println("POST")
	}
	req, err := http.NewRequest(method, path, strings.NewReader(""))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return string(body), nil
}