package api

import (
	"fmt"
	tgbotapi "github.com/huantt/telegram-bot-api/v5"
	"github.com/huantt/telegram-bot-test/api/webhook"
	"log"
	"net/http"
	"os"
)

func Init() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("%+v", err)
	}
	webhookHandler := webhook.NewHandler(webhook.NewService(bot))
	http.HandleFunc("/webhook/updates", webhookHandler.Handle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
