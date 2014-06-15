package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"io"
	"os/user"
	"log"
	"io/ioutil"
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
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}

	dat, err := ioutil.ReadFile(usr.HomeDir +"/.goGithubApiSample")
	check(err)
	fmt.Print(string(dat))

	// TODO: 起動引数ありでトークン更新/なしで保存しているトークンを使う
	access_token := string(dat)
	access_token = strings.Replace(access_token, "\n", "", -1)

	url3 := "https://api.github.com/issues?state=open&filter=all&access_token=" + access_token
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func httpRequestLog(urlStr string) string {
	resp, err1 := http.Get(urlStr)
	if err1 != nil {
		// handle error
		log.Fatal("err:", err1.Error())
	}
	defer resp.Body.Close()
	//defer resp.Header.Close()

	for key, value := range resp.Header {
		fmt.Println(key, value)
	}

//	fmt.Println("")

	body, _ := ioutil.ReadAll(resp.Body)
	responseBody := string(body[:])
//	fmt.Println("body :", responseBody)

//	fmt.Println("")

	return responseBody
}
