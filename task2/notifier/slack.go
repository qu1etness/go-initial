package notifier

import (
	"errors"
	"log/slog"
	"math/rand"
)

type SlackClient struct {
	apiKey string
}

func NewSlackClient(apiKey string) *SlackClient {
	return &SlackClient{
		apiKey: apiKey,
	}
}

func (s *SlackClient) SendToChannel(channel string, message string) error {

	slog.Info("Sending message to channel", slog.String("message", message), slog.String("channel", channel))
	if rand.Intn(100)%2 == 0 {
		return errors.New("failed to send message")
	}

	return nil
}

func (s *SlackClient) SendDirectMessage(userId string, message string) error {
	slog.Info("Sending direct message", slog.String("message", message), slog.String("userId", userId))

	if rand.Intn(100)%2 == 0 {
		return errors.New("failed to send message")
	}
	return nil
}

func (s *SlackClient) GetBotStatus() bool {
	return rand.Intn(100)%2 == 0
}
