package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

func main() {

	var _url string = "http://www.jktree.com/disease/article/bfd6.html"

	fmt.Println("start ...")

	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:56.0) Gecko/20100101 Firefox/56.0"
	c.Limit(&colly.LimitRule{DomainGlob: "www.jktree.com", Parallelism: 1})

	c.OnHTML("article", func(e *colly.HTMLElement) {
		fmt.Println("#### ", e.ChildText("h1"), "####")
		fmt.Println(e.ChildText("#articleBody"))
	})

	// Find and visit all links
	n := 0
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if strings.Contains(url, ".html") && strings.Contains(url, "article") {
			fmt.Printf("=== %d %s\n", n, url)
			time.Sleep(time.Second * 5)
			e.Request.Visit(url)
		}
		n = n + 1
	})

	//c.OnHTML("body", func(e *colly.HTMLElement) {
	//fmt.Println(e.DOM.Find("article").Html())
	//})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("-> GET ", r.URL)
	})

	c.Visit(_url)

}
