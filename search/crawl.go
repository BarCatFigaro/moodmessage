package search

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
)

// Extender defines the extension methods required by the crawler
type Extender struct {
	*gocrawl.DefaultExtender
}

// Visit does the scraping work
func (e *Extender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	// use goquery here
	fmt.Printf("Visited: %s\n", ctx.URL())
	return nil, true
}

// Filter only crawled reddit pages
func (e *Extender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited
}

// NewSpider returns a new crawler
func NewSpider() *gocrawl.Crawler {
	ext := &Extender{&gocrawl.DefaultExtender{}}
	opts := gocrawl.NewOptions(ext)
	opts.RobotUserAgent = "moodmessage"
	opts.UserAgent = "Mozilla/5.0 (compatible; moodmessage/1.0)"
	opts.CrawlDelay = 2 * time.Second
	opts.LogFlags = gocrawl.LogAll
	opts.SameHostOnly = false
	opts.MaxVisits = 50

	c := gocrawl.NewCrawlerWithOptions(opts)

	return c
}

// RunSpider starts the web crawler at seed
func RunSpider(c *gocrawl.Crawler, seed string) {
	c.Run(seed)
}
