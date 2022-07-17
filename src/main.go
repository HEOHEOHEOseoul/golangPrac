package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var statusError = errors.New("status err")

func main() {
	var results = map[string]string{}
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitUrl(url, c)

	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}

}

func hitUrl(url string, c chan requestResult) {

	resp, err := http.Get(url)
	status := "ok"
	if err != nil || resp.StatusCode >= 400 {
		status = "failed"
	}
	c <- requestResult{url: url, status: status}

}

// func main() {
// 	c := make(chan string)
// 	people := [2]string{"hi", "bye"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 		fmt.Println(<-c)
// 	}

// }
// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 3)
// 	c <- person + " is sexy"
// }
