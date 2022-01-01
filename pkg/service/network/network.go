package network

import (
	"log"

	"github.com/frostyjet/telegram-bot/pkg/scrapper"
)

type Network struct {
}

func NewNetwork() *Network {
	return &Network{}
}

func (n *Network) SearchKeyword(query string) (description string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	url := n.BuildUrl(query)

	log.Print("URL: ", url)

	scrapper.DownloadContent(url, query)
	description = scrapper.ParseFile(query)

	return description, err
}

func (n *Network) BuildUrl(query string) string {
	return "https://en.wikipedia.org/wiki/" + query
}
