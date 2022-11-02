package handlers

import (
	"interview/domain/constants"
	"interview/domain/entities"
	"interview/internal/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type InternalOrdersHandler struct {
	repository repositories.OrdersRepository
}

func NewInternalOrdersHandler(repository repositories.OrdersRepository) Handler {
	return &InternalOrdersHandler{repository}
}

func (h *InternalOrdersHandler) RegisterRoutes(router *echo.Group, mws map[string]echo.MiddlewareFunc) {
	internalRoutes := router.Group("/internal", mws[constants.Admin])
	internalRoutes.GET(constants.OrdersPath, h.List)
	internalRoutes.POST(constants.OrdersPath, h.Create)
	internalRoutes.GET(constants.OrderWithIDPath, h.Retrieve)
	internalRoutes.PUT(constants.OrderWithIDPath, h.Update)
	internalRoutes.DELETE(constants.OrderWithIDPath, h.Destroy)
}

func (h *InternalOrdersHandler) List(c echo.Context) error {
	orders, err := h.repository.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"orders": orders}

	return c.JSON(http.StatusOK, response)
}

func (h *InternalOrdersHandler) Create(c echo.Context) error {
	order := &entities.Order{}
	if err := c.Bind(order); err != nil {
		return err
	}

	err := h.repository.Create(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, order)
}

func (h *InternalOrdersHandler) Retrieve(c echo.Context) error {
	var id uint64

	err := echo.PathParamsBinder(c).Uint64("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order, err := h.repository.Retrieve(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func (h *InternalOrdersHandler) Update(c echo.Context) error {
	var id uint

	err := echo.PathParamsBinder(c).Uint("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	order := &entities.Order{}

	if err := c.Bind(order); err != nil {
		return err
	}

	order.ID = id

	err = h.repository.Update(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *InternalOrdersHandler) Destroy(c echo.Context) error {
	var id uint64

	err := echo.PathParamsBinder(c).Uint64("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	err = h.repository.Destroy(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
