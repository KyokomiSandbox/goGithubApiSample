package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println("")

	httpRequestLog("https://api.github.com/users/octocat/orgs")

	url2 := "https://api.github.com/repos/kyokomi/cocostudio_book/issues"
	httpRequestLog(url2)
}

func httpRequestLog(urlStr string) {
	resp, err1 := http.Get(urlStr)
	if err1 != nil {
		// handle error
		fmt.Printf("err:", err1.Error())
	}
	defer resp.Body.Close()
	//defer resp.Header.Close()

	for key, value := range resp.Header {
		fmt.Println(key, value)
	}

	fmt.Println("")

	body, _ := ioutil.ReadAll(resp.Body)
	responseBody := string(body[:])
	fmt.Printf("body :", responseBody)

	fmt.Println("")
}
