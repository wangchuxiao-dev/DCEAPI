package DCEAPI

import (
	"net/http"
	"strings"
	"io/ioutil"
	_ "net/url"
)

type Exchange struct {
	Name string
	Apikey string
	Secret string
	Password string
	LimitRate bool
	Debug bool
}

func BaseRequest(method string, path string, body string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func RequestBuilder() {
	
}
