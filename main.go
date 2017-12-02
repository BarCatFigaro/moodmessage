package main

import "fmt"
import "./twPull"
import "github.com/saintpete/twilio-go"
import "time"

func main() {
    now := time.Now()
    messages := []twilio.Message{}
    newmessages := []string{}
    for {
        messages, newmessages = twPull.GetMessages(now,messages)
        for _,msg := range newmessages {
            fmt.Println(msg)
        }
    }

}
