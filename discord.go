package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func ConnectToDiscord() {
	dg, err := discordgo.New("Bot " + os.Getenv("Token"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	switch m.Content {
	case "!test":
		s.ChannelMessageSend(m.ChannelID, "accepted")

	case "!latency":
		message := ("Your latency with discord is : " + s.HeartbeatLatency().String())
		s.ChannelMessageSend(m.ChannelID, message)

	case "!join":
		s.Identify.Intents = discordgo.IntentsGuildVoiceStates
		v, err := s.ChannelVoiceJoin(m.GuildID, "923235701309464620", true, false)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		v.Speaking(true)

	}
}
