package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {

	// Request the HTML page.
	res, err := http.Get("https://www.gmo.jp/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	var bandlist []string
	doc.Find(".js-tab-block ul .list-news-unit").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		bandlist = append(bandlist, band)
	})
	//fmt.Printf("Review : %s \n", bandlist)
	// uniq news
	m := make(map[string]bool)
	uniq := []string{}
	for _, ele := range bandlist {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
			fmt.Printf("Review : %s \n", ele)
		}
	}
}
func main() {
	ExampleScrape()
}
