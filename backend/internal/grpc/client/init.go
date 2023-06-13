package client

import (
	"backend/conf/settings"
	"backend/internal/grpc/client/match"
	"backend/internal/grpc/client/snake"
)

func Init() {
	match.Init(settings.Conf)
	snake.Init(settings.Conf)
}
