package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanrafli14/CatalogAPI/internal/service"
	"github.com/ivanrafli14/CatalogAPI/pkg/middleware"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRes(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndPoint() {

	auth := r.router.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/login", r.Login)

	category := r.router.Group("/category")
	category.GET("/", r.GetAllCategory)

	product := r.router.Group("/product")
	product.Use(r.middleware.Authentication)
	product.POST("/", r.CreateProduct)
	product.GET("/", r.GetProduct)

	file := r.router.Group("/files")
	file.Use(r.middleware.Authentication)
	file.POST("/upload", r.UploadProductImage)

}

func (r *Rest) Run() {
	r.router.Run(":8080")
}
