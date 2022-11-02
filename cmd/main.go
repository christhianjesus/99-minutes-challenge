package main

import (
	"fmt"
	"interview/cmd/config"
	"interview/domain/constants"
	"interview/domain/entities"
	"interview/internal/handlers"
	"interview/internal/middlewares"
	"interview/internal/repositories"

	"github.com/joeshaw/envdecode"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := config.Config{}

	if err := envdecode.Decode(&conf); err != nil {
		panic(fmt.Errorf("Cannot read from env: %w", err))
	}

	db, err := gorm.Open(postgres.Open(conf.DSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	s := echo.New()

	// Router
	router := s.Group(constants.PrefixPath)

	// Configure middleware with the custom claims type
	JWTConf := middleware.JWTConfig{
		Claims:     &entities.JwtCustomClaims{},
		SigningKey: []byte(conf.JWTSecret),
	}

	// Middlewares
	mws := map[string]echo.MiddlewareFunc{
		constants.User:  middleware.JWTWithConfig(JWTConf),
		constants.Admin: middlewares.AdminAuth(conf.AdminToken),
	}

	setupHandlers(&ServerInstances{
		conf,
		db,
		router,
		mws,
	})

	s.Logger.Debug(
		s.Start(conf.Addr()),
	)
}

type ServerInstances struct {
	conf   config.Config
	db     *gorm.DB
	router *echo.Group
	mws    map[string]echo.MiddlewareFunc
}

func setupHandlers(i *ServerInstances) {
	authHandler := handlers.NewAuthHandler(i.conf.JWTSecret)
	ordersRepository := repositories.NewOrdersRepository(i.db)
	internalOrdersHandler := handlers.NewOrdersHandler(ordersRepository)

	authHandler.RegisterRoutes(i.router, i.mws)
	internalOrdersHandler.RegisterRoutes(i.router, i.mws)
}
