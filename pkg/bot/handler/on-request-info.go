package handler

import (
	"log"
	"strings"

	"github.com/frostyjet/telegram-bot/entities"
	"github.com/frostyjet/telegram-bot/pkg/helpers"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (h *Handler) OnRequestInfo(m *tb.Message) {
	db := h.Services.Db
	net := h.Services.Network

	query := strings.TrimSpace(m.Payload)

	query = helpers.RemoveSpecialChars(query)
	query = helpers.PrepareQuery(query)

	quote, ok := db.Get(query)
	var response string

	if !ok || quote.Description == "" {
		log.Print("No local information found for query: ", query)

		description, err := net.SearchKeyword(query)

		if err == nil {
			quote = entities.Quote{
				Description: description,
			}

			db.Set(query, quote)
			db.Persist()
		}
	}

	if quote.Description != "" {
		response = "This is what I found about <b>" + m.Payload + "<b>\n\n" + quote.Description
	} else {
		response = "Sorry, I couldn't find any information about " + m.Payload
	}

	h.Bot.Send(m.Sender, response)
}
