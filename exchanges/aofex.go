package exchanges

import (
	_ "crypto/sha1"
	"github.com/PythonohtyP1900/DCEAPI"
	_ "net/url"
	"crypto/rand" //真随机
    "fmt"
    "math/big"
)

const Path string = "https://aofex.com/"
const CNPath string = "https://aofex.co/"
const SpotPath string = "https://openapi.aofex.com/"
const SwapPath string = "https://openapi-contract.aofex.com/"
const DebugSpotPath string = ""
const DebugSwapPath string = ""
const nonceInt = "78"

type Aofex struct {
	Path string
	SpotPath string
	SwapPath string
	Exchange *DCEAPI.Exchange
}

func NewAofex(debug bool, secret, apikey string) *Aofex {
	var spotPath string
	var swapPath string
	if debug == true {
		spotPath = DebugSpotPath
		swapPath = DebugSwapPath
	} else {
		spotPath = SpotPath
		swapPath = SwapPath
	}
	aofex := &Aofex{
		Path: Path,
		SpotPath: spotPath,
		SwapPath: swapPath,
		Exchange: &DCEAPI.Exchange{
			Name: "AOFEX",
			Secret: secret,
			Apikey: apikey,
			Debug: debug,
		},
	}
	return aofex
}

func sign(apikey, secret string, data map[string]string) {
	nonce, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(nonce)
	tmp := []string{apikey, secret, nonceInt}
	for k, v := range data {
		tmp = append(tmp, k+"="+v)
	}
	var hashString string
	for _, v := range tmp {
		hashString += v
	}
	t := sha1.New();
	io.WriteString(t,data);
}

func getHeader() {
	
}

func (aofex *Aofex) FetchMarkets() (string, error) {
	route := "openApi/market/symbols"
	url := aofex.SpotPath + route
	res, err := DCEAPI.BaseRequest("GET", url, "")
	return res, err
}

func (aofex *Aofex) FetchTrades() (string, error) {
	route := "openApi/market/trade"
	url := aofex.SpotPath + route
	res, err := DCEAPI.BaseRequest("GET", url, "")
	return res, err
}

func (aofex *Aofex) LimitBuyOrder(symbol string, amount, price float64) {
	
}

func (aofex *Aofex) LimitSellOrder(symbol string, amount, price float64) {

}

func (aofex *Aofex) FetchBalance() {

}
