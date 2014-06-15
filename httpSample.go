package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"io"
	"log"
)

type Issue struct {
	Url string `json:"url"`
	Number int `json:"number"`
	Title string `json:"title"`
}

type ResponseData struct {
	Issues []Issue
}

func main() {
	// TODO: 起動引数ありでトークン更新/なしで保存しているトークンを使う
	/**
		- [ ] 起動引数指定できるようにする
		- [ ] 文字列を暗号化？とかして保存する。どこに？ ~/.{appName}とか？
		- [ ] 暗号化した文字列をファイル？を読み込む（ ~/.{appName}とか？）
	 */
	access_token := ""

	url3 := "https://api.github.com/issues?access_token=" + access_token
	fmt.Println("--- " + url3 + " ---")
	responseJson := httpRequestLog(url3)

	dec := json.NewDecoder(strings.NewReader(responseJson))
	for {
		var res []Issue
		if err := dec.Decode(&res); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for _, i := range res {
			fmt.Printf("issue %d: [%s] -> %s\n", i.Number, i.Title, i.Url)
		}
	}
}

func httpRequestLog(urlStr string) string {
	resp, err1 := http.Get(urlStr)
	if err1 != nil {
		// handle error
		fmt.Println("err:", err1.Error())
	}
	defer resp.Body.Close()
	//defer resp.Header.Close()

	for key, value := range resp.Header {
		fmt.Println(key, value)
	}

	fmt.Println("")

	body, _ := ioutil.ReadAll(resp.Body)
	responseBody := string(body[:])
	fmt.Println("body :", responseBody)

	fmt.Println("")

	return responseBody
}
