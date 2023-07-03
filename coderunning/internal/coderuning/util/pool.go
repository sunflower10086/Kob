package util

import (
	"coderunning/internal/models"
	"context"
	"fmt"
	"sync"
)

var (
	Bots = make(chan models.Bot, 10000)
	mu   = sync.Mutex{}
)

func AddBot(userId int32, botCode, input string) {
	mu.Lock()
	defer mu.Unlock()

	Bots <- models.Bot{UserId: userId, BotCode: botCode, Input: input}
}

func Run(ctx context.Context) {
	for {
		select {
		case bot := <-Bots:
			TaskNum <- struct{}{}
			err := consume(ctx, bot)
			if err != nil {
				fmt.Printf(err.Error())
			}
		}
	}
}
