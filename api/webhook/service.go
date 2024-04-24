package webhook

import (
	"context"
	"fmt"
	tgbotapi "github.com/huantt/telegram-bot-api/v5"
	"log/slog"
	"time"
)

type Service struct {
	bot *tgbotapi.BotAPI
}

func NewService(bot *tgbotapi.BotAPI) *Service {
	return &Service{bot: bot}
}

func (s *Service) Handle(ctx context.Context, update *tgbotapi.Update) error {
	if update != nil && update.Message != nil {
		slog.Info(fmt.Sprintf("Latency: %s", time.Since(update.Message.Time())))
		_, err := s.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Hello"))
		return err
	}
	return nil
}
