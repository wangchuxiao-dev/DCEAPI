package DCEAPI

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	"time"
	"errors"
)

type Exchange struct {
	Name string
	Apikey string
	Secret string
	Password string
	LimitRate bool
	Debug bool
}

func BuildRequestUrl(url string, params map[string]string) (string) {
	if params != nil {
		url += "?"
	}
	for k, v := range params {
		url += (k+"="+v+"&")
	}
	return url
}

// 封装基础请求, url已经通过参数builder
func HttpRequest(method, path, body string, headers map[string]string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 1,
	}
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla")
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HttpStatusCode:%d ,Desc:%s", resp.StatusCode, string(resBody)))
	}
	return resBody, nil
}
