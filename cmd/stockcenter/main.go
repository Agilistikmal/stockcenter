package main

import (
	"github.com/agilistikmal/stockcenter/internal/stockcenter/delivery/discord"
	"github.com/agilistikmal/stockcenter/internal/stockcenter/infrastructure/config"
)

func main() {
	config.Load()
	discordClient := discord.New()
	discordClient.Run()
}
