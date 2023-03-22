package sender

import (
	"fmt"
	"go_cli/pkg/logger"
	"io"
	"net/http"
	"os"
	"strings"
)

type SlackSender struct {
	botToken  string
	channelID string
}

func NewSlackSender() *SlackSender {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if botToken == "" {
		logger.Error("SLACK_BOT_TOKEN environment variable is not set")
		return nil
	}

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		logger.Error("SLACK_CHANNEL_ID environment variable is not set")
		return nil
	}
	return &SlackSender{botToken: botToken, channelID: channelID}
}

func (s *SlackSender) Configure() error {
	panic("implement me")
}

func (s *SlackSender) Send(message string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+s.botToken)
	req.Header.Set("Content-Type", "application/json")

	body := fmt.Sprintf(`{"channel": "%s", "text": "%s"}`, s.channelID, message)
	req.Body = io.NopCloser(strings.NewReader(body))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}

func (s *SlackSender) GetName() string {
	return "Slack"
}
