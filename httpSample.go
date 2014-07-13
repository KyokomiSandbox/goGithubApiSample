package main

import (
	"fmt"
	"flag"
	"net/http"
	"strings"
	"os/user"
	"log"
	"os"
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

	var accessTokenFlag string

	/* register flag name and shorthand name */
	accessTokenMessage := "Hear your github account Access Token."
	flag.StringVar(&accessTokenFlag, "access-token", "", accessTokenMessage)
	flag.StringVar(&accessTokenFlag, "a"           , "", accessTokenMessage)
	flag.Parse()

	var access_token string

	if accessTokenFlag != "" {
		log.Println("access-token param ok.")

		access_token = accessTokenFlag

		byteData := []byte(access_token)
		writeFilePath := usr.HomeDir +"/.goGithubApiSample"
		err := ioutil.WriteFile(writeFilePath, byteData, os.ModePerm)
		if err != nil {
			log.Fatal("access token write error.")
			check(err)
		} else {
			log.Println(writeFilePath + " save access-token.")
		}

	} else {
		log.Println("not access-token param read saveFile.")

		dat, err := ioutil.ReadFile(usr.HomeDir +"/.goGithubApiSample")
		if err != nil {
			log.Fatal("not access token error.")
			check(err)
		}
		fmt.Print(string(dat))

		accessTokenSaveFile := string(dat)
		accessTokenSaveFile = strings.Replace(accessTokenSaveFile, "\n", "", -1)
		if accessTokenSaveFile != "" {
			access_token = accessTokenSaveFile
		} else {
			panic("not access token error.")
		}
	}
	fmt.Println("access_token =", access_token)

	t := &oauth.Transport {
		Token: &oauth.Token{AccessToken: access_token},
	}
	client := github.NewClient(t.Client())

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List("", nil)
	check(err)
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
