package exchanges

import (
	"crypto/sha1"
	_ "net/url"
	_ "crypto/rand" 
	"fmt"
	"sort"
	_ "math/big"
	"io"
	"time"
	"github.com/PythonohtyP1900/DCEAPI"
)

const (
	Path string = "https://aofex.com/"
	CNPath string = "https://aofex.co/"
	SpotPath string = "https://openapi.aofex.co"
	SwapPath string = "https://openapi-contract.aofex.co"
	DebugSpotPath string = ""
	DebugSwapPath string = ""
)

type BaseResponse struct {
	Errno int `json errno`
	Errmsg string `json errmsg`
	result interface{}
}

type Aofex struct {
	Path string
	SpotPath string
	SwapPath string
	Exchange *DCEAPI.Exchange
}

func NewAofex(secret, apikey string) *Aofex {
	aofex := &Aofex{
		Path: Path,
		SpotPath: SpotPath,
		SwapPath: SwapPath,
		Exchange: &DCEAPI.Exchange{
			Name: "AOFEX",
			Secret: secret,
			Apikey: apikey,
		},
	}
	return aofex
}

func DebugNewAofex(secret, apikey string) *Aofex {
	aofex := &Aofex{
		Path: Path,
		SpotPath: DebugSpotPath,
		SwapPath: DebugSwapPath,
		Exchange: &DCEAPI.Exchange{
			Name: "AOFEX",
			Secret: secret,
			Apikey: apikey,
		},
	}
	return aofex
}

func sign(apikey, secret, nonce string, data map[string]string) string {
	tmp := sort.StringSlice{apikey, secret, nonce}
	for k, v := range data {
		tmp = append(tmp, k+"="+v)
	}
	fmt.Println(tmp)
	sort.Sort(tmp)
	var hashString string
	for _, v := range tmp {
		hashString += v
	}
	t := sha1.New();
	io.WriteString(t,hashString);
	return fmt.Sprintf("%x",t.Sum(nil));
}

func generateNonce() string {
	nonce := int(time.Now().Unix())
	
	return nonce.String() + "_" + "sadx1"
}

func generateHeader(apikey, secret string, params map[string]string) map[string]string {
	nonce := generateNonce()
	fmt.Println(nonce)
	return map[string]string{
		"Nonce": nonce,
		"Token": apikey,
		"Signature": sign(apikey, secret, nonce, params),
	}
}

func (aofex *Aofex) aofexRequestPublic(method, path string, params map[string]string) (string, error) {
	return "1", nil
}

func (aofex *Aofex) aofexRequestPrivate(method, path string, params map[string]string) (string, error) {
	headers := generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, params)
	res_string, err := DCEAPI.BaseRequest(method, path, "", headers)
	return res_string, err
}

func (aofex *Aofex) FetchBalance() (string, error) {
	res_string, err := aofex.aofexRequestPrivate("GET", aofex.SpotPath+"/openApi/wallet/list?show_all=1", map[string]string{"show_all":"1"})
	return res_string, err
}


func (aofex *Aofex) FetchMarkets() (string, error) {
	
	route := "openApi/market/symbols"
	url := aofex.SpotPath + route
	
	res, err := aofex.aofexRequestPublic("GET", url, map[string]string{})
	return res, err
}

func (aofex *Aofex) FetchTrades() (string, error) {
	route := "openApi/market/trade"
	url := aofex.SpotPath + route
	res, err := DCEAPI.BaseRequest("GET", url, "", nil)
	return res, err
}

// func (aofex *Aofex) LimitBuyOrder(symbol string, amount, price float64) {
// 	params := map[string]interface{}{"symbol": symbol, "amount":amount, "price":price}
// 	url := "testurl"
// 	// res, err := aofex.aofexRequestPrivate("POST", url, params)
		
// }

// func (aofex *Aofex) LimitSellOrder(symbol string, amount, price float64) {

// }

