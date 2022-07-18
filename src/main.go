package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://upbit.com/service_center/notice"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	fmt.Println(doc)
	doc.Find("#UpbitLayout > div.subMain > div > section.ty02 > article").Each(func(i int, s *goquery.Selection) {
		// test, test2 := s.Find("#bodySec_2").Html()
		fmt.Println(s.Html())
	})
	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("request status:", res.StatusCode)
	}
}
