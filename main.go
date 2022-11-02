package minuts

import (
	"fmt"
	"interview/config"
	"interview/constants"
	"interview/handlers"
	"interview/middleware"
	"interview/repositories"

	"github.com/joeshaw/envdecode"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.Config{}

	if err := envdecode.Decode(&conf); err != nil {
		panic(fmt.Errorf("Cannot read from env: %w", err))
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
		mws,
		router,
	})

	s.Logger.Fatal(
		s.Start(conf.Addr()),
	)
}

type ServerInstances struct {
	conf   config.Config
	mws    map[string]echo.MiddlewareFunc
	router *echo.Group
}

func setupHandlers(i *ServerInstances) {
	ordersRepository := repositories.NewOrdersRepository()
	internalOrdersHandler := handlers.NewInternalOrdersHandler(ordersRepository)

	internalOrdersHandler.RegisterRoutes(i.router, i.mws)
}
