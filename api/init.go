package api

import (
	"fmt"
	"github.com/huantt/telegram-bot-test/api/webhook"
	"log"
	"net/http"
	"os"
)

func Init() {
	webhookHandler := webhook.NewHandler(webhook.NewService())
	http.HandleFunc("/webhook/updates", webhookHandler.Handle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
