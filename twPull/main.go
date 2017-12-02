package twPull

import "log"
//import "fmt"
import "github.com/saintpete/twilio-go"
import "context"
import "net/url"
import "time"
//import "io/ioutil"
import "os"
import "encoding/json"
import "fmt"

type Config struct {
    Sid string `json:"sid"`
    Token string `json:"token"`
}

func messageExists(a twilio.Message, list []twilio.Message) bool {
    for _,b := range list {
        if b.Sid == a.Sid {
            return true
        }
    }
    return false
}

func GetMessages(start time.Time, out []twilio.Message) ([]twilio.Message, []string) {
    now := time.Now()
    file, e := os.Open("./twiliosecret.json")
    defer file.Close()
    if e != nil {
        fmt.Println(e)
        os.Exit(1)
    }

    var config Config
    jsonParser := json.NewDecoder(file)
    jsonParser.Decode(&config)


    sid := config.Sid
    token := config.Token
    
    client := twilio.NewClient(sid, token, nil)

    iter := client.Messages.GetMessagesInRange(start, now, url.Values{})

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    messagestrings := []string{}
    for {
        page, err := iter.Next(ctx)
        if err == twilio.NoMoreResults {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        for _, message := range page.Messages {
 //           fmt.Printf("%d: %s (%s)\n", i,  message.Body, message.To)
            if message.To == "+16139098924" {
                if !messageExists(*message,out) {
                    out = append(out, *message)
                    messagestrings = append(messagestrings,message.Body)
                }
            }
        }
    }
    return out,messagestrings

}

