package main

import (
	"flag"
	"fmt"
	"gophercises/link"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gravitational/trace"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	check(err)
	defer resp.Body.Close()

	// need to add domain to paths -> /some-path
	// need to deal with fragment or mailto links
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	// use the link parsing package from previous exercise
	links, _ := link.Parse(resp.Body)
	var hrefs []string
	for _, ll := range links {
		switch {
		case strings.HasPrefix(ll.Href, "/"):
			hrefs = append(hrefs, base+ll.Href)
		case strings.HasPrefix(ll.Href, "http"):
			hrefs = append(hrefs, ll.Href)
		}
	}
	for _, href := range hrefs {
		fmt.Println(href)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(trace.DebugReport(err))
	}
}

// 1. GET the webpage
// 2. parse all the links on the page
// 3. build proper urls with our links
// 4. filter out any links with different domain
// 5. find all pages (BFS)
// 6. print outl xml
