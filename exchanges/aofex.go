package exchanges

import (
	"crypto/sha1"
	"github.com/PythonohtyP1900/DCEAPI"
	_ "net/url"
	"crypto/rand" 
    "fmt"
	"math/big"
	"io"
)

const Path string = "https://aofex.com/"
const CNPath string = "https://aofex.co/"
const SpotPath string = "https://openapi.aofex.com/"
const SwapPath string = "https://openapi-contract.aofex.com/"
const DebugSpotPath string = ""
const DebugSwapPath string = ""

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

func sign(apikey, secret string, data map[string]string) {
	nonce, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(nonce)
	tmp := []string{apikey, secret}
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
	data := map[string]interface{}{"symbol": symbol, "amount":amount, "price":price}

	
}

func (aofex *Aofex) LimitSellOrder(symbol string, amount, price float64) {

}

func (aofex *Aofex) FetchBalance() {

}
