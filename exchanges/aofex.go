package exchanges

import (
	"crypto/sha1"
	"math/rand" 
	"fmt"
	"sort"
	"strconv"
	"io"
	"encoding/json"
	"time"

	"github.com/PythonohtyP1900/DCEAPI"
)

const (
	PATH string = "https://aofex.com/"
	CNPATH string = "https://aofex.co/"
	SPOTPATH string = "https://openapi.aofex.co"
	SWAPPATH string = "https://openapi-contract.aofex.co"
)

type Aofex struct {
	Path string
	SpotPath string
	SwapPath string
	Exchange *DCEAPI.Exchange
}		

func NewAofex(secret, apiKey string) *Aofex {
	aofex := &Aofex{
		Path: PATH,
		SpotPath: SPOTPATH,
		SwapPath: SWAPPATH,
		Exchange: &DCEAPI.Exchange{
			Name: "AOFEX",
			Secret: secret,
			ApiKey: apiKey,
		},
	}
	return aofex
}

// 签名
func (aofex *Aofex) sign(apikey, secret, nonce string, data map[string]string) string {
	tmp := sort.StringSlice{apikey, secret, nonce}
	for k, v := range data {
		tmp = append(tmp, k+"="+v)
	}
	sort.Sort(tmp)
	var hashString string
	for _, v := range tmp {
		hashString += v
	}
	t := sha1.New();
	io.WriteString(t, hashString);
	return fmt.Sprintf("%x",t.Sum(nil));
}

// 生成随机字符串
func (aofex *Aofex) getRandomStr() string {
	rand.Seed(time.Now().UnixNano())
	str := "abcdefghijklmnopqrstuvwxyz"
	var randStr string
	for i:=0;i<4;i++ {
		randInt := rand.Intn(len(str)-1)
		randStr += string(str[randInt])
	}
	return randStr + strconv.Itoa(rand.Intn(9))
}

// 时间戳+随机字符串
func (aofex *Aofex) generateNonce() string {
	ts := strconv.FormatInt(time.Now().Unix(),10)
	return ts + "_" + aofex.getRandomStr()
}

// 生成请求头
func (aofex *Aofex) generateHeader(apikey, secret string, params map[string]string) map[string]string {
	nonce := aofex.generateNonce()
	return map[string]string{
		"Nonce": nonce,
		"Token": apikey,
		"Signature": aofex.sign(apikey, secret, nonce, params),
	}
}

func (aofex *Aofex) GetExchangeName() string {
	return "AOFEX"
}

type AofexBaseResponse struct {
	Errno int 
	ErrMsg string 
}

type aofexResponse interface {
	hasError() error
}

func (baseResponse *AofexBaseResponse) hasError() error {
	var err error
	switch baseResponse.Errno {
		case 0:
			err = nil
		case 20504:
			err = &DCEAPI.BalanceError{ErrCode:20504, ErrMsg:baseResponse.ErrMsg}
		case 20501, 20502:
			err = &DCEAPI.SymbolError{ErrCode:20501, ErrMsg:baseResponse.ErrMsg}
		case 20506, 20522, 20521:
			err = &DCEAPI.InvalidSignatureError{ErrCode:baseResponse.Errno, ErrMsg:baseResponse.ErrMsg}
		case 20510, 20511, 20512, 20513, 20514:
			err = &DCEAPI.OrderLimitError{ErrCode:baseResponse.Errno, ErrMsg:baseResponse.ErrMsg}
		case 20515:
			err = &DCEAPI.OrderNotFound{ErrCode:20515, ErrMsg:baseResponse.ErrMsg}
		case 20516:
			err = &DCEAPI.OrderStateError{ErrCode:20516, ErrMsg:baseResponse.ErrMsg}
		default:
			err = &DCEAPI.ExchangeError{ErrCode:baseResponse.Errno, ErrMsg:baseResponse.ErrMsg}
	}
	return err
}

func (aofex *Aofex) request(method, path string, params, body, headers map[string]string, model aofexResponse) error {
	var err error
	path = DCEAPI.BuildRequestUrl(path, params)
	bodyStr, err := DCEAPI.BuildRequestBody(body)
	if err != nil {
		return err
	}
	if headers != nil {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}
	res, err := DCEAPI.HttpRequest(method, path, bodyStr, headers)
	// fmt.Println(string(res))
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(res, model)
	if err != nil {
		return err
	}
	err = model.hasError()
	return err
}
