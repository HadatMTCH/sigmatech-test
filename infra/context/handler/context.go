package handler

import (
	"base-api/app/http/handlers/template"
	"base-api/infra/context/service"
)

type HandlerContext struct {
	TemplateHandler template.Template
}

// initServiceCtx for contextService
func InitHandlerContext(serviceContext *service.ServiceContext) *HandlerContext {
	return &HandlerContext{
		TemplateHandler: template.New(serviceContext),
	}
}
