package service

import (
	"base-api/app/services/template"
	"base-api/config"
	"base-api/infra/context/module"
	"base-api/infra/context/repository"
	"base-api/infra/middleware"
	"base-api/infra/s3"

	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	TemplateService template.Template
	S3Service       s3.S3Interface
	JWTService      middleware.JWTInterface
}

// initServiceCtx for contextService
func InitServiceContext(repositoryContext *repository.RepositoryContext, moduleContext *module.ModuleContext, config *config.Config, redisClient *redis.Client) *ServiceContext {
	return &ServiceContext{
		TemplateService: template.New(repositoryContext),
		JWTService:      middleware.NewJWT(&config.JWTConfig, redisClient),
		S3Service:       s3.NewS3Configuration(&config.S3),
	}
}
