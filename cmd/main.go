package main

import (
	"fmt"
	"interview/cmd/config"
	"interview/domain/constants"
	"interview/internal/handlers"
	"interview/internal/middleware"
	"interview/internal/repositories"

	"github.com/joeshaw/envdecode"
	"github.com/labstack/echo/v4"
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

	// Middlewares
	mws := map[string]echo.MiddlewareFunc{
		constants.User:  middleware.UserAuth(),
		constants.Admin: middleware.AdminAuth(conf.AdminToken),
	}

	setupHandlers(&ServerInstances{
		conf,
		db,
		mws,
		router,
	})

	s.Logger.Debug(
		s.Start(conf.Addr()),
	)
}

type ServerInstances struct {
	conf   config.Config
	db     *gorm.DB
	mws    map[string]echo.MiddlewareFunc
	router *echo.Group
}

func setupHandlers(i *ServerInstances) {
	ordersRepository := repositories.NewOrdersRepository(i.db)
	internalOrdersHandler := handlers.NewInternalOrdersHandler(ordersRepository)

	internalOrdersHandler.RegisterRoutes(i.router, i.mws)
}
