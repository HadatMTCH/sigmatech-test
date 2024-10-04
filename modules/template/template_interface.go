package template

import (
	"base-api/infra/context/repository"
	"base-api/infra/fcm"
	"base-api/infra/websocket"

	"github.com/go-redis/redis/v8"
)

type Template interface {
}

func New(ctx *repository.RepositoryContext, redis *redis.Client, wsClient websocket.WebsocketInterface, fcm fcm.FCMInterface) Template {
	return &template{
		ctx,
		redis,
		wsClient,
		fcm,
	}
}
