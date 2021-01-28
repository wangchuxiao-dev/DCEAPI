package main

import (
	"fmt"
	"encoding/json"
)

func BuildRequestBody(params interface{}) (string, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	jsonBody := string(data)
	return jsonBody, nil
}


func main()  {

	url, err := BuildRequestBody("ss")
	fmt.Println(url, err)
}