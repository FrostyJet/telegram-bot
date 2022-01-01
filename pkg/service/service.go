package service

import (
	"github.com/frostyjet/telegram-bot/pkg/service/cache"
	"github.com/frostyjet/telegram-bot/pkg/service/network"
)

type Service struct {
	Db      *cache.Cache
	Network *network.Network
}

func NewService(config Service) *Service {
	return &Service{
		Db:      config.Db,
		Network: config.Network,
	}
}
