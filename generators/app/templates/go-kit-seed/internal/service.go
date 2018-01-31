package internal

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

// Service - interface
type Service interface {
	SampleEndPoint(ctx context.Context, req []Data) ([]Data, error)
}

type service struct {
	DB *gorm.DB
}

// ServiceConfig - config stuct for service
type ServiceConfig struct {
	Logger log.Logger
	DB     *gorm.DB
}

func initService(config ServiceConfig) Service {
	//Create / update database model if nessasary
	// READMORE -  http://jinzhu.me/gorm/database.html#migration
	config.DB.AutoMigrate(&Data{})
	return service{config.DB}
}

// NewService - Service Constructor
func NewService(config ServiceConfig) Service {
	var s Service

	// NOTE - Expect additional middlewares here.
	s = initService(config)
	// s = newPermissionMiddleware(s)
	// s = newLoggingMiddleware(s, config)

	return s
}

func (s service) SampleEndPoint(ctx context.Context, req []Data) ([]Data, error) {
	return []Data{}, errors.New("NOT IMPLEMENTED")
}
