package main

import (
	"fmt"
	"net/http"
	"strings"
	"os/user"
	"log"
	"io/ioutil"
	"github.com/google/go-github/github"
	"code.google.com/p/goauth2/oauth"
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

	// TODO: 起動引数ありでトークン更新/なしで保存しているトークンを使うようにする
	access_token := string(dat)
	access_token = strings.Replace(access_token, "\n", "", -1)

	t := &oauth.Transport {
		Token: &oauth.Token{AccessToken: access_token},
	}
	client := github.NewClient(t.Client())

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List("", nil)
	for idx, repo := range repos {
		fmt.Println(idx, *repo.Name)
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
