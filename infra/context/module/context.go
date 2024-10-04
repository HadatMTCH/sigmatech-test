package module

import (
	"base-api/config"
	"base-api/infra/context/repository"
	"base-api/infra/fcm"
	"base-api/infra/redis"
	"base-api/infra/websocket"
	"base-api/modules/template"
	"context"
	"log"
)

type ModuleContext struct {
	Template template.Template
}

// initModuleContext for 3rd party modules
func InitializeModuleContext(ctx *repository.RepositoryContext, cfg *config.Config) *ModuleContext {
	redisServer := redis.NewRedisServer(&cfg.Redis)
	wsClient := websocket.NewWebsocket(&cfg.Server)
	rdb, err := redisServer.Connect(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	notification := fcm.NewFCM(&cfg.FCM)

	templateMod := template.New(ctx, rdb, wsClient, notification)

	return &ModuleContext{
		Template: templateMod,
	}
}
