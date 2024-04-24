package webhook

import (
	"context"
	"fmt"
	tgbotapi "github.com/huantt/telegram-bot-api/v5"
	"log/slog"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Handle(ctx context.Context, update *tgbotapi.Update) error {
	if update != nil && update.Message != nil {
		slog.Info(fmt.Sprintf("Latency: %s", time.Since(update.Message.Time())))
	}
	return nil
}
