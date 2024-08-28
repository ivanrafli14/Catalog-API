package main

import (
	"github.com/ivanrafli14/CatalogAPI/internal/handler/rest"
	"github.com/ivanrafli14/CatalogAPI/internal/repository"
	"github.com/ivanrafli14/CatalogAPI/internal/service"
	"github.com/ivanrafli14/CatalogAPI/pkg/bcrypt"
	"github.com/ivanrafli14/CatalogAPI/pkg/cloudinary"
	"github.com/ivanrafli14/CatalogAPI/pkg/config"
	"github.com/ivanrafli14/CatalogAPI/pkg/database/postgres"
	"github.com/ivanrafli14/CatalogAPI/pkg/database/redis"
	"github.com/ivanrafli14/CatalogAPI/pkg/jwt"
	"github.com/ivanrafli14/CatalogAPI/pkg/meilisearch"
	"github.com/ivanrafli14/CatalogAPI/pkg/middleware"
)

func main() {
	config.LoadEnv()
	jwt := jwt.Init()
	bcrypt := bcrypt.Init()
	redis := redis.Init()
	cloudinay := cloudinary.Init()
	meilisearch := meilisearch.Init() 
	
	db := postgres.ConnectDB()
	postgres.Migrate(db)
	postgres.SeedDB(db)
	
	repo := repository.NewRepository(db)
	service:= service.NewService(service.InitParam{Repository: repo, Bcrypt: bcrypt, JWTAuth: jwt, Redis: redis, Cloudinary: cloudinay, Meilisearch: meilisearch})
	middleware:= middleware.Init(jwt, service,redis)
	rest :=rest.NewRes(service, middleware)
	rest.MountEndPoint()
	rest.Run()

	
}