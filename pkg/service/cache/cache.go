package cache

import (
	"encoding/json"
	"os"

	"github.com/frostyjet/telegram-bot/entities"
)

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
	storageContents, _ := os.ReadFile("resources/cache.json")
	json.Unmarshal(storageContents, &c.data)
}

func (c *Cache) Persist() {
	storageContents, _ := json.Marshal(c.data)
	os.WriteFile("resources/cache.json", storageContents, 0644)
}
