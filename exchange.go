package DCEAPI

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Exchange struct {
	Name      string
	ApiKey    string
	Secret    string
	Password  string
	LimitRate bool
	Debug     bool
}

func BuildRequestUrl(path string, params map[string]string) string {
	value := url.Values{}
	for k, v := range params {
		value.Add(k, v)
	}
	if params != nil {
		path += "?"
	}
	return path + value.Encode()
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
	// fmt.Println(req.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, HttpError{resp.StatusCode, string(resBody)}
	}
	return resBody, nil
}
