package event

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(client *discordgo.Session, event *discordgo.Ready) {
	fmt.Printf("Discord bot %s is now online\n", event.User.Username)
	client.UpdateWatchStatus(0, "Stock List")
}
