package search

import (
	"fmt"
	"log"
	"time"

	"github.com/turnage/graw"

	"github.com/turnage/graw/reddit"
)

// NewBot creates a new reddit bot instance
func NewBot() reddit.Bot {
	bot, err := reddit.NewBotFromAgentFile("search/agentfile.template", 5*time.Second)
	if err != nil {
		log.Fatalf("could not start bot: %v\n", err)
	}
	return bot
}

// RunBot runs the bot
func RunBot(bot reddit.Bot, cfg graw.Config) {
	stop, wait, err := graw.Run(&Announcer{}, bot, cfg)

	_ = stop

	if err != nil {
		fmt.Printf("graw run starting error: %v\n", err)
	}

	if err := wait(); err != nil {
		fmt.Printf("graw run encountered an error %v\n", err)
	}
}

// Announcer is a handler for Run
type Announcer struct{}

// Post is called when there is a new post
func (a *Announcer) Post(post *reddit.Post) error {
	fmt.Printf("%s posted: %s and created %d\n", post.Author, post.SelfText, post.CreatedUTC)
	return nil
}

// Comment implements Comment handler
func (a *Announcer) Comment(post *reddit.Comment) error {
	fmt.Printf("%s posted: %s", post.Author, post.Body)
	return nil
}
