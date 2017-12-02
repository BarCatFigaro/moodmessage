package search

import "github.com/turnage/graw"

// NewConfig creates a new graw configuration
func NewConfig() graw.Config {
	cfg := graw.Config{
		Subreddits: []string{"self", "moodmessage"},
	}

	return cfg
}
