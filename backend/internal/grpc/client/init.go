package client

import (
	"backend/conf/settings"
	"backend/internal/grpc/client/match"
)

func Init() {
	match.Init(settings.Conf)
}
