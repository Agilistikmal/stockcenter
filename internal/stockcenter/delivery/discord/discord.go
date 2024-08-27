package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/agilistikmal/stockcenter/internal/stockcenter/delivery/discord/event"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

type DiscordHandler struct {
	client *discordgo.Session
}

func New() *DiscordHandler {
	client, err := discordgo.New("Bot " + viper.GetString("discord.token"))
	if err != nil {
		log.Fatal(err)
	}

	client.Identify.Intents = discordgo.IntentDirectMessages

	return &DiscordHandler{
		client: client,
	}
}

func (h *DiscordHandler) Run() {
	h.SetupHandler()

	err := h.client.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Discord Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	h.client.Close()
}

func (h *DiscordHandler) SetupHandler() {
	h.client.AddHandler(event.Ready)
}
