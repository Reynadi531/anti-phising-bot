package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Reynadi531/anti-phsing-discord/events"
	"github.com/bwmarrin/discordgo"
)

var (
	Token = os.Getenv("TOKEN")
)

func main() {
	dg, err := discordgo.New(fmt.Sprintf("Bot %s", Token))
	if err != nil {
		fmt.Println("Error creating session: ", err)
	}

	dg.AddHandler(events.MessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("The bot is running")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
