package infra

import (
	"base-api/config"
	"base-api/infra/context/handler"
	"base-api/infra/context/module"
	"base-api/infra/context/repository"
	"base-api/infra/context/service"
	"base-api/infra/db"
	"base-api/infra/middleware"
	"base-api/infra/redis"
	"context"
	"log"
)

type InfraContextInterface interface {
	Config() config.Config
	Handler() *handler.HandlerContext
	Middleware() *multipleMiddleware
}

type multipleMiddleware struct {
	TokenMiddleware middleware.TokenMiddlewareInterface
}

type infraContext struct {
	config          config.Config
	tokenMiddleware middleware.TokenMiddlewareInterface
	handlerContext  *handler.HandlerContext
}

func New() InfraContextInterface {
	// initial config
	ctx := context.Background()
	cfg := config.InitConfig()

	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := db.Open(&cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}

	// initial redis server
	redisServer := redis.NewRedisServer(&cfg.Redis)
	rdb, err := redisServer.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// init repo context
	repositoryContext := repository.InitializeRepositoryContext(db, &cfg.S3)

	// init module context
	moduleContext := module.InitializeModuleContext(repositoryContext, &cfg)

	// init service context
	serviceContext := service.InitServiceContext(repositoryContext, moduleContext, &cfg, rdb)

	// Initialize Handler Context
	handlercontext := handler.InitHandlerContext(serviceContext)

	// init middleware ctx
	tokenMiddleware := middleware.NewTokenMiddleware(serviceContext.JWTService)
	return &infraContext{
		handlerContext:  handlercontext,
		config:          cfg,
		tokenMiddleware: tokenMiddleware,
	}
}

func (i infraContext) Config() config.Config {
	return i.config
}

func (i infraContext) Handler() *handler.HandlerContext {
	return i.handlerContext
}

func (i infraContext) Middleware() *multipleMiddleware {
	return &multipleMiddleware{
		TokenMiddleware: i.tokenMiddleware,
	}
}
