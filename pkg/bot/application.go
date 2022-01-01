package bot

import (
	"log"

	"github.com/frostyjet/telegram-bot/pkg/service"
	"github.com/frostyjet/telegram-bot/pkg/service/cache"
	"github.com/frostyjet/telegram-bot/pkg/service/network"
)

var db *cache.Cache
var net *network.Network

type ApplicationConfig struct {
	Token string
}

type Application struct {
	Config ApplicationConfig
}

func NewApplication(token string) *Application {
	return &Application{
		Config: ApplicationConfig{Token: token},
	}
}

func (app *Application) Run() {
	log.Println("Creating cache")
	db = cache.NewCache()

	log.Println("Loading cache from last persist")
	db.Load()

	log.Println("Initializing search network")
	net = network.NewNetwork()

	app.InitRouter(&service.Service{
		Db:      db,
		Network: net,
	})
}
