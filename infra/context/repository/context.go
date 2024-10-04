package repository

import (
	"base-api/app/repositories/template"
	"base-api/config"
	"base-api/infra/db"
)

type RepositoryContext struct {
	TemplateRepository template.Template
	DB                 *db.DB
}

func InitializeRepositoryContext(db *db.DB, config *config.S3Configuration) *RepositoryContext {
	templateRepository := template.New(db)

	return &RepositoryContext{
		DB:                 db,
		TemplateRepository: templateRepository,
	}
}
