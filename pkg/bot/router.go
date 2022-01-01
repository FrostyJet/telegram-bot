package bot

import (
	"log"
	"strings"
	"time"

	"github.com/frostyjet/telegram-bot/pkg/bot/handler"
	"github.com/frostyjet/telegram-bot/pkg/service"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (app *Application) InitRouter(s *service.Service) {

	log.Print("Configuring bot...")
	tbBot, err := tb.NewBot(tb.Settings{
		Token:  app.Config.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	h := handler.NewHandler(s, tbBot)

	tbBot.Handle(tb.OnText, func(m *tb.Message) {
		text := m.Text

		if strings.HasPrefix(text, "who is") {
			m.Payload = strings.Replace(text, "who is", "", 1)
			h.OnRequestInfo(m)
			return
		}

		tbBot.Send(m.Sender, "Hello, "+m.Sender.FirstName+"!")
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print("Starting bot...")
	tbBot.Start()
}
