 package main

import (
	_ "github.com/PythonohtyP1900/DCEAPI"
	"github.com/PythonohtyP1900/DCEAPI/exchanges"
	"fmt"
)

func main() {
	token := "4bec6394e490aca7acaae197379824d3"
	secret := "k51r7mii94jlebhk4ahq"
	aofex := exchanges.NewAofex(secret, token)
	a, err := aofex.FetchBalance()
	fmt.Println(a, err)
}