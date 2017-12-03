package search

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
)

// Only enqueue the root and comments/posts
var rxOk = regexp.MustCompile(`https?:\/\/(www\.)?reddit\.com\/r\/UBC(\/comments.*)?$`)

// Extender defines the extension methods required by the crawler
type Extender struct {
	*gocrawl.DefaultExtender
}

// Visit does the scraping work
func (e *Extender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	// use goquery here
	fmt.Printf("Visited: %s\n", ctx.URL())
	body := ""
	doc.Find(".usertext-body .may-blank-within .md-container").Each(func(index int, item *goquery.Selection) {
		body = item.Find("p").Text()
	})

	fmt.Println(body)
	return nil, true
}

// Filter only crawled reddit pages
func (e *Extender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited && rxOk.MatchString(ctx.NormalizedURL().String())
}

// NewSpider returns a new crawler
func NewSpider() *gocrawl.Crawler {
	ext := &Extender{&gocrawl.DefaultExtender{}}
	opts := gocrawl.NewOptions(ext)
	opts.RobotUserAgent = "moodmessage"
	opts.UserAgent = "Mozilla/5.0 (compatible; moodmessage/1.0)"
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogAll
	opts.MaxVisits = 3

	c := gocrawl.NewCrawlerWithOptions(opts)

	return c
}

// RunSpider starts the web crawler at seed
func RunSpider(c *gocrawl.Crawler, seed string) {
	c.Run(seed)
}
