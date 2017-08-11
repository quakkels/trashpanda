package main

import (
	"fmt"
	"github.com/turnage/graw/reddit"
)

func main() {
	bot, err := reddit.NewBotFromAgentFile("trashpanda.agent", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle. ", err)
		return
	}

	harvest, err := bot.Listing("/r/golang", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/golang. ", err)
		return
	}

	bot.SendMessage("quakkels", "You're being msg'd", harvest.Posts[0].Title)
	fmt.Println("You've been msg'd about: ", harvest.Posts[0].Title)
}
