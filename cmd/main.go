package main

import (
	"log"
	"my-web-scraper/cmd/store"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var cars []store.CarPageDetails 


func main() {
	response, err := http.Get("https://jiji.co.ke/lavington/cars/mercedes-benz-gle-class-gle-400d-4matic-2020-white-6XVeDUZtUTTG07BSIwScMSQ.html?page=1&pos=1&cur_pos=1&ads_per_page=20&ads_count=20&lid=r-HIpnFySDmUzsMFV4")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	var records []string

	records = append(records, "Title")

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	cars = make([]store.CarPageDetails, 0)
	detail := make([]store.Detail, 0)

	doc.Find("div.qa-advert-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("div.b-advert-title-inner").Text()
		s.Find("div.b-advert-item-details-collapser").Each(func( i int, k *goquery.Selection) {
			key := k.Find("div.b-advert-attribute__key").Text()
            value := k.Find("div.b-advert-attribute__value").Text()
			keys := strings.Split(key, "\n")
			values := strings.Split(value, "\n")
			for i := range keys {
				keys[i] = strings.TrimSpace(keys[i])
                values[i] = strings.TrimSpace(values[i])
				if keys[i] == "" && values[i] == "" {
                    continue
                }
				d := store.Detail{Key: keys[i], Value: values[i]}
				detail = append(detail, d)
			}
		})
		cars = append(cars, store.CarPageDetails{Title: title, Details: detail})
		detail = make([]store.Detail, 0)
	})
	store.SaveToJSON(cars)
}
