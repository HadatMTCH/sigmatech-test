package template

import (
	"base-api/infra/context/repository"
	"base-api/infra/fcm"

	"base-api/infra/websocket"

	"github.com/go-redis/redis/v8"
)

type template struct {
	*repository.RepositoryContext
	redis    *redis.Client
	wsClient websocket.WebsocketInterface
	fcm.FCMInterface
}
