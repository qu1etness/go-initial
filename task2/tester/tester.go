package tester

import "log/slog"

type TesterService struct {
	LastMessage string
}

func NewTesterService() *TesterService {
	return &TesterService{
		LastMessage: "",
	}
}

func (t *TesterService) SendToChannel(channel string, message string) error {
	t.LastMessage = message
	slog.Info("CPU Load is critically high!", slog.String("channel", channel), slog.String("message", message))

	return nil
}
