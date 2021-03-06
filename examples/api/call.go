package main

import (
	"fmt"
	"github.com/xu42/youzan-sdk-go"
)

func main() {

	resp, err := youzan.GenSelfToken("CLIENT_ID", "CLIENT_SECRET", "110")
	if err != nil {
		fmt.Println(err)
		return
	}

	params := map[string]string{
		"page_no":   "1",
		"page_size": "10",
	}

	result, err := youzan.Call(resp.AccessToken, "youzan.scrm.customer.search", "3.1.0", params)
	fmt.Println(result.Success, result.Result, result.Error, err)
}
