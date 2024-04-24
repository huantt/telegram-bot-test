package api

import (
	"github.com/huantt/telegram-bot-test/api/webhook"
	"log"
	"net/http"
)

func Init() {
	webhookHandler := webhook.NewHandler(webhook.NewService())
	http.HandleFunc("/webhook/updates", webhookHandler.Handle)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
