package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	discordToken := os.Getenv("DISCORD_TOKEN")

	if discordToken == "" {
		log.Fatal("cannot find discord token")
	}

	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}

	dg.AddHandler(messageCreate)

  dg.Identify.Intents = discordgo.IntentGuildMessages

  err = dg.Open()
  if err != nil{
    log.Fatal("error opening connection,",err)
  }

  log.Println("Bot is now running. Press CTRL-C to exit.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc,syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
  <-sc


  // close discord session
  dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Author.Bot {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
	}

}
