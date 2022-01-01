package bot

import (
	"log"
	"strings"
	"time"

	"github.com/frostyjet/telegram-bot/pkg/bot/handler"
	"github.com/frostyjet/telegram-bot/pkg/service"
	tb "gopkg.in/tucnak/telebot.v2"
)

const welcomeMsg = `I'm a bot and I am still in development.<br>
Please, ask me about stuff you would like to know by 
starting your sentence with <em>who is</em> or <em>what is</em>, 
for example <em>who is Mariah Carey</em>`

func (app *Application) InitRouter(s *service.Service) {

	log.Print("Configuring bot...")
	tbBot, err := tb.NewBot(tb.Settings{
		Token:     app.Config.Token,
		Poller:    &tb.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tb.ModeHTML,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	h := handler.NewHandler(s, tbBot)

	tbBot.Handle(tb.OnText, func(m *tb.Message) {
		text := m.Text

		if strings.HasPrefix(text, "who is") || strings.HasPrefix(text, "what is") {
			text = strings.Replace(text, "who is", "", 1)
			m.Payload = strings.Replace(text, "what is", "", 1)
			h.OnRequestInfo(m)
			return
		}

		tbBot.Send(m.Sender, "Hello, "+m.Sender.FirstName+"!\n"+welcomeMsg)
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print("Starting bot...")
	tbBot.Start()
}
