package main

import (
	"github.com/joho/godotenv"
	"github.com/shaned24/crabbot-discord/crabbot"
	"log"
	"os"
	"os/signal"
	"syscall"

	"toughcrab.com/halo/autocode"
	"toughcrab.com/halo/bot"
)

const (
	AutoCodeUrl = "https://halo.api.stdlib.com/infinite@0.3.3/stats/service-record/multiplayer/"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not find or load .env file, skipping...")
	}

	discordToken := os.Getenv("DISCORD_TOKEN")
	autoCodeToken := os.Getenv("AUTOCODE_TOKEN")

	client := autocode.NewAutoCodeClient(AutoCodeUrl, autoCodeToken)

	botInstance, err := crabbot.NewBot(
		discordToken,
		"!",
		bot.NewServiceRecordMultiplayer(client),
	)

	if err != nil {
		log.Println("error creating Bot session,", err)
	}

	defer botInstance.Close()

	err = botInstance.Start()
	if err != nil {
		log.Printf("Couldn't start the bot: %v", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
