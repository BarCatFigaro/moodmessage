package main

import "fmt"
import "./twPull"
import "./stringsem"
import "github.com/saintpete/twilio-go"
import "time"
import (
	"github.com/barcatfigaro/moodmessage/search"
)

func main() {
    now := time.Now()
	spider := search.NewSpider()
	search.RunSpider(spider, "https://reddit.com/")

    messages := []twilio.Message{}
    newmessages := []string{}
    for {
        messages, newmessages = twPull.GetMessages(now,messages)
        for _,msg := range newmessages {
            fmt.Println(msg)
            if stringsem.IsGood(msg){
                fmt.Println("that was a good message")
            } else {
                fmt.Println("that was bad message")
            }
        }
    }
}

