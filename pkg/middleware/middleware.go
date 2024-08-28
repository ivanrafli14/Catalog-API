package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanrafli14/CatalogAPI/internal/service"
	"github.com/ivanrafli14/CatalogAPI/pkg/database/redis"
	"github.com/ivanrafli14/CatalogAPI/pkg/jwt"
)

type Interface interface {
	Authentication(c *gin.Context)
}

type middleware struct {
	jwtAuth jwt.Interface
	redis redis.Interface
	service *service.Service
}

func Init(jwtAuth jwt.Interface, service *service.Service, redis redis.Interface) Interface {
	return &middleware{
		jwtAuth: jwtAuth,
		redis: redis,
		service: service,
	}
}