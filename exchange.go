package DCEAPI

import (
	"net/http"
	"strings"
	"io/ioutil"
	_ "net/url"
	"encoding/json"
	"fmt"
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

func BuildRequestBody(params interface{}) (string, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	jsonBody := string(data)
	return jsonBody, nil
}

// 封装基础请求, url已经通过参数builder
func HttpRequest(method, path, body string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	fmt.Println(body)
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla")
	req.Header.Add("content-type", "application/json")
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
	fmt.Println(string(resBody))
	return resBody, nil
}
