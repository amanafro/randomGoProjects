package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func webscrapping() {
	defer timer("main")()

	c := colly.NewCollector()
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		h.Request.Visit(h.Attr("href"))
	})

	c.OnHTML("h1", func(h *colly.HTMLElement) {
		h.Request.Visit(h.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	var website string
	fmt.Println("Enter website")
	fmt.Scan(&website)

	c.Visit(website)
}
