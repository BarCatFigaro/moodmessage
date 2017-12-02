package main

import (
	"github.com/barcatfigaro/moodmessage/search"
)

func main() {

	/*
		bot := search.NewBot()
		cfg := search.NewConfig()
		search.RunBot(bot, cfg)
	*/

	spider := search.NewSpider()
	search.RunSpider(spider, "https://reddit.com/")

	/*
			now := time.Now()
			messages := []twilio.Message{}
			newmessages := []string{}
			for {
				messages, newmessages = twPull.GetMessages(now, messages)
				for _, msg := range newmessages {
					fmt.Println(msg)
				}
		    }
	*/
}
