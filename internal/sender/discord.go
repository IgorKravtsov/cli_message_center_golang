package sender

import (
	"github.com/bwmarrin/discordgo"
	"go_cli/pkg/interfaceManager"
	"go_cli/pkg/logger"
	"os"
)

type DiscordSender struct {
	botToken  string
	channelID string
}

func NewDiscordSender() *DiscordSender {
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if botToken == "" {
		logger.Error("DISCORD_BOT_TOKEN environment variable is not set")
		return nil
	}
	channelID := os.Getenv("DISCORD_CHANNEL_ID")
	if channelID == "" {
		channelID = "1087040420883742763"
	}
	return &DiscordSender{botToken, channelID}
}

func (ds *DiscordSender) Configure() error {
	interfaceManager.Clear()

	interfaceManager.SetText("Enter the Discord bot token: ")
	token, err := interfaceManager.GetEnteredText()
	if err != nil {
		return err
	}
	ds.botToken = token

	interfaceManager.SetText("Enter the Discord channel Id: ")
	channelID, err := interfaceManager.GetEnteredText()
	if err != nil {
		return err
	}
	ds.channelID = channelID

	return nil
}

func (ds *DiscordSender) Send(msg string) error {
	dg, err := discordgo.New("Bot " + ds.botToken)
	if err != nil {
		return err
	}
	defer dg.Close()

	_, err = dg.ChannelMessageSend(ds.channelID, msg)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DiscordSender) GetName() string {
	return "Discord"
}
