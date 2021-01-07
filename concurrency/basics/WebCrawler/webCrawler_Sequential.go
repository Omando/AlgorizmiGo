package WebCrawler

import (
	"AlgorizmiGo/collections/sequential"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/html"		// Use cmd: go get golang.org/x/net
)

// Web crawling using breadth-first traversal
func Crawl(topUrl string) {
	// This map is used to ensure that we do not crawl the same url twice
	processed := make(map[string]bool)

	// Used in breadth traversal
	var queue = sequential.CreateQueue()
	queue.Enqueue(topUrl)

	for !queue.IsEmpty() {
		item, err  := queue.Dequeue()
		if err != nil {
			fmt.Printf("Failed to queue item: %v", err)
			continue
		}

		// Crawl this url and return a list of new urls to crawl
		url := item.(string)
		links := crawlUrl(url)
		processed[url] = true
		for _, link := range links {
			// Queue this link if not processed before
			if !processed[link] {
				queue.Enqueue(link)
			}
		}
	}
}

// A thin wrapper around extract to get and return links found in a URL
func crawlUrl(url string) (links []string) {
	fmt.Println(url)
	links, err := extract(url)
	if err != nil {
		log.Fatal(err)
	}
	return links
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document
// Overall sequence of steps:
//	1. Call http.Get
//  2. Check for error
//  3. Check response status
//	4. Parse the body using html.Parse
//  5. Perform depth-first traversal to extract href values
func extract(url string) (values []string, err error) {
	// Make the http request
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// When err is nil, resp always contains a non-nil resp.Body
	// Always close resp.Body when done reading from it
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("error loading %s: %s", url, res.Status)
	}

	// Parse body of response
	doc, err := html.Parse(res.Body)
	if err != nil {
		res.Body.Close()
		return nil, fmt.Errorf("Error parsing body as html: %v", err)
	}

	// Create a function value that takes and parses an html node; get the node's attributes
	// and parse the value of any href attribute
	var hrefVisitor = func(n *html.Node) {
		// Process attributes
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attribute := range n.Attr {
				// Do we have an "href" attribute?
				if attribute.Key != "href" {
					continue
				}

				// We have an href attribute. Get its value and parse it
				hrefValue, err := res.Request.URL.Parse(attribute.Val)
				if err != nil {
					continue		// ignore bad urls
				}

				values = append(values, hrefValue.String())
			}
		}
	}

	// Do a depth first traversal on all nodes found in doc. For each node, use hrefVisitor
	// to extract href values and populate the links slice
	visitNode(doc, hrefVisitor, nil)
	return values, nil
}

func visitNode(doc *html.Node, pre, post func(n *html.Node)) {
	// Pre-visit node if required
	if pre != nil {
		pre(doc)
	}

	// Depth-first traversal
	for child := doc.FirstChild; child != nil; child = child.NextSibling {
		visitNode(child, pre, post)
	}

	// Post-visit node if required
	if post != nil {
		post(nil)
	}
}
