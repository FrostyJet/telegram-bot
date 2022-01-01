package handler

import (
	"github.com/frostyjet/telegram-bot/pkg/service"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Handler struct {
	Services *service.Service
	Bot      *tb.Bot
}

func NewHandler(s *service.Service, b *tb.Bot) *Handler {
	return &Handler{
		Services: s,
		Bot:      b,
	}
}
