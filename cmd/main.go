package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, err := http.Get("https://jiji.co.ke/")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("Title %d: %s\n", i+1, title)
	})

}
