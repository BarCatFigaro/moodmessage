package search

import (
	"fmt"
	"log"
	"time"

	"github.com/barcatfigaro/moodmessage/stringsem"

	"github.com/turnage/graw"

	"github.com/turnage/graw/reddit"
)

//var messages *[]string
var channel chan []string
var m_messages []string
// NewBot creates a new reddit bot instance
func NewBot(c chan []string) reddit.Bot {
	bot, err := reddit.NewBotFromAgentFile("search/agentfile.template", time.Second)
//	messages = arr
	if err != nil {
		log.Fatalf("could not start bot: %v\n", err)
	}
    channel = c
	return bot
}

// RunBot runs the bot
func RunBot(bot reddit.Bot, cfg graw.Config) {
    stop, wait, err := graw.Run(&Announcer{bot: bot}, bot, cfg)

	_ = stop

	if err != nil {
		fmt.Printf("graw run starting error: %v\n", err)
	}



	if err := wait(); err != nil {
		fmt.Printf("graw run encountered an error %v\n", err)
	}
}

// Announcer is a handler for Run
type Announcer struct{ bot reddit.Bot }

// Post is called when there is a new post
func (a *Announcer) Post(post *reddit.Post) error {
    fmt.Println(post.Title)

    select {
    case m_messages = <-channel:
        fmt.Println("recievedmsg")
    default:
        fmt.Println("nomsgrecieved")
    }
    m_post := post.Title + "\n" + post.SelfText
    fmt.Println(m_post)
	if !stringsem.IsGood(m_post) {
		return a.bot.SendMessage(
			post.Author,
			fmt.Sprintf("MoodMessage: %s", m_post),
            m_messages[len(m_messages)-1],
		)
	}

	return nil
}

// Comment implements Comment handler
func (a *Announcer) Comment(post *reddit.Comment) error {
	fmt.Printf("COMMENT: %s posted: %s", post.Author, post.Body)
	return nil
}
