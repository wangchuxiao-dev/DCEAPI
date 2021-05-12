package exchanges

import (
	"github.com/PythonohtyP1900/DCEAPI"

	"encoding/json"
	"strings"
	"net/url"
	"time"
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

const (
	HUOBI_PATH string = "https://api.huobi.pro"
)

type Huobi struct {
	Path string
	SpotPath string
	SwapPath string
	SpotID int
	Exchange *DCEAPI.Exchange
}

func NewHuobi(secret, apiKey string) *Huobi {
	huobi := &Huobi{
		Path: HUOBI_PATH,
		Exchange: &DCEAPI.Exchange{
			Name: "HUOBI",
			ApiKey: apiKey,
			Secret: secret,
		},
	}
	account, err := huobi.GetAccount()
	if err != nil {
		panic("huobi初始化失败")
	}
	huobi.SpotID = account["spot"]
	return huobi
}

type huobiBaseResponse struct {
	Ts int `json:"ts"`
	Status string `json:"status"`
	ErrCode string `json:"err-code"`
	ErrMsg string `json:"err-msg"`
}

func (baseResponse huobiBaseResponse) hasError() error {
	var err error
	if baseResponse.Status != "ok" {
		err = &DCEAPI.ExchangeError{ErrMsg:baseResponse.ErrCode, ErrCode:500}
	}
	return err
}

type huobiResponse interface {
	hasError() error
}


func (huobi *Huobi) sign(host, method, path, secret string, params map[string]string) string {
	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteString("\n")
	host = "api.huobi.pro"
	sb.WriteString(host)
	sb.WriteString("\n")
	sb.WriteString(path)
	sb.WriteString("\n")
	value := url.Values{}
	for k, v := range params {
		value.Add(k, v)
	}
	sb.WriteString(value.Encode())
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(sb.String()))
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return result
}

func (huobi *Huobi) generateSignature(host, method, path, secret, apiKey string, params map[string]string) map[string]string {
	signMap := map[string]string{
		"AccessKeyId": apiKey,
		"SignatureMethod": "HmacSHA256",
		"SignatureVersion": "2",
		"Timestamp": time.Now().UTC().Format("2006-01-02T15:04:05"),
	}
	for k, v := range params {
		signMap[k] = v
	}
	signature := huobi.sign(host, method, path, secret, signMap)
	signMap["Signature"] = signature
	return signMap
}

func (huobi Huobi) GetExchangeName() string {
	return "HuobiPro"
}

func (huobi Huobi) request(method, path string, params, body map[string]string, model huobiResponse) error {
	var err error
	path = DCEAPI.BuildRequestUrl(path, params)
	headers := map[string]string{}
	if body != nil {
		headers["Content-Type"] = "application/json"
	}
	bytesData, err := json.Marshal(body)
	if err != nil {
		return err
	}
	res, err := DCEAPI.HttpRequest(method, path, string(bytesData), headers)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	err = json.Unmarshal(res, model)
	if err != nil {
		return err
	}
	err = model.hasError()
	return err
}

func (huobi Huobi) GetAccount() (map[string]int, error) {
	url := "/v1/account/accounts"
	type AccountResponse struct {
		huobiBaseResponse
		Data []struct{
			Id int `json:"id"`
			Type string `json:"type"` 
		} `json:"data"`
	}
	accountResponse := &AccountResponse{}
	signMap := huobi.generateSignature(huobi.Path, "GET", url, huobi.Exchange.Secret, huobi.Exchange.ApiKey, nil)
	err := huobi.request("GET", huobi.Path+url, signMap, nil, accountResponse)
	account := map[string]int{}
	for _, v := range accountResponse.Data {
		account["spot"] = v.Id
	}
	return account, err
} 

func (huobi Huobi) SymbolFormatConversion(symbol string) string {
	temp := strings.Split(strings.ToLower(symbol), "/")
	newSymbol := temp[0] + temp[1]
	return newSymbol
}