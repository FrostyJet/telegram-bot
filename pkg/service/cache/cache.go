package cache

import (
	"encoding/json"
	"os"

	"github.com/frostyjet/telegram-bot/entities"
)

const cachePath = "resources/cache.json"

type Cache struct {
	data map[string]entities.Quote
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]entities.Quote),
	}
}

func (c *Cache) Set(query string, quote entities.Quote) {
	c.data[query] = quote
}

func (c *Cache) Get(query string) (entities.Quote, bool) {
	quote, ok := c.data[query]

	return quote, ok
}

func (c *Cache) Load() {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	storageContents, err := os.ReadFile(basePath + "/" + cachePath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(storageContents, &c.data)
}

func (c *Cache) Persist() {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	storageContents, _ := json.Marshal(c.data)
	os.WriteFile(basePath+"/"+cachePath, storageContents, 0644)
}
