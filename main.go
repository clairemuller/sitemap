package main

import (
	"flag"
	"fmt"
	"gophercises/link"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gravitational/trace"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()
	pages := get(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}
}

// get takes in the url given by the user, makes a GET request,
// determines the base url, and passes the response body and url to our hrefs function,
// returning a slice of hrefs
func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	check(err)
	defer resp.Body.Close()

	// want to use the request url from the response
	// in case the site that the user passes in redirects
	reqURL := resp.Request.URL
	// baseURL will be just the domain
	// in case the user passed in url with paths
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	links := hrefs(resp.Body, base)
	return filter(base, links)
}

// hrefs takes in the response body and base url,
// uses link.Parse to get all the links,
// then checks the links, adding the domain if needed
// and returns a slice of hrefs (links)
func hrefs(r io.Reader, base string) []string {
	// use the link parsing package from previous exercise
	// returns a slice of links
	links, _ := link.Parse(r)
	var hrefsSlice []string
	// need to add domain to paths -> /some-path
	// need to deal with fragment or mailto links
	for _, ll := range links {
		switch {
		case strings.HasPrefix(ll.Href, "/"):
			hrefsSlice = append(hrefsSlice, base+ll.Href)
		case strings.HasPrefix(ll.Href, "http"):
			hrefsSlice = append(hrefsSlice, ll.Href)
		}
	}
	return hrefsSlice
}

// filter out any links that don't have the base url
func filter(base string, links []string) []string {
	var linksSlice []string
	for _, link := range links {
		if strings.HasPrefix(link, base) {
			linksSlice = append(linksSlice, link)
		}
	}
	return linksSlice
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
// 6. print outl xml - not doing this
