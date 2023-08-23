package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://oa.api2d.net/dashboard/billing/credit_grants"
	method := "GET"

	payload := strings.NewReader(`{
    "model": "text-davinci-edit-001",
    "instruction": "请修改文本中的拼写错误",
    "input": "What tim is it"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer fk213819-LV6bFDfyzGfztX2jXhBqdTRR8OvkLaCQ")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	summary := CreditSummary{}

	// 解析 JSON 数据到结构体
	err = json.Unmarshal([]byte(body), &summary)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return
	}

	// 提取 total_available 的值
	totalAvailable := summary.TotalAvailable
	fmt.Println("当前余额:", totalAvailable)
}

type CreditSummary struct {
	TotalAvailable int `json:"total_available"`
}
