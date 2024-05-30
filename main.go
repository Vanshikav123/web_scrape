package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func fetchHTML(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func scrapeDoc(url string) {
	doc, err := fetchHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	// Check if any elements are found
	found := false
	doc.Find("article.IFHyqb").Each(func(i int, s *goquery.Selection) {
		found = true
		headline := s.Text()
		fmt.Printf("%d: %s\n", i+1, headline)
	})

	if !found {
		fmt.Println("No headlines found with the given selector.")
	}
}

func main() {
	url := "https://news.google.com/home?hl=en-IN&gl=IN&ceid=IN:en"
	fmt.Println("Fetching headlines from:", url)
	scrapeDoc(url)
}
