package main

import (
	"fmt"
	"time"

	"github.com/barcatfigaro/moodmessage/search"
	"github.com/barcatfigaro/moodmessage/stringsem"
	"github.com/barcatfigaro/moodmessage/twPull"
	twilio "github.com/saintpete/twilio-go"
)

func main() {
	now := time.Now()
	spider := search.NewSpider()
	search.RunSpider(spider, "https://www.reddit.com/r/UBC/")

	messages := []twilio.Message{}
	newmessages := []string{}
	for {
		messages, newmessages = twPull.GetMessages(now, messages)
		for _, msg := range newmessages {
			fmt.Println(msg)
			if stringsem.IsGood(msg) {
				fmt.Println("that was a good message")
			} else {
				fmt.Println("that was bad message")
			}
		}
	}
}
