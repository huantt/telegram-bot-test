package webhook

import (
	"context"
	"encoding/json"
	tgbotapi "github.com/huantt/telegram-bot-api/v5"
	"io"
	"log/slog"
	"net/http"
)

type Handler struct {
	srv IService
}

type IService interface {
	Handle(ctx context.Context, update *tgbotapi.Update) error
}

func NewHandler(srv IService) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var update tgbotapi.Update
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
	err = h.srv.Handle(r.Context(), &update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := io.WriteString(w, "OK"); err != nil {
		slog.Error(err.Error())
	}
}
