package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gravitational/trace"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	check(err)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
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
