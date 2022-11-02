package handlers

import (
	"interview/constants"
	"interview/entities"
	"interview/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrdersHandler struct {
	repository repositories.OrdersRepository
}

func NewOrdersHandler(repository repositories.OrdersRepository) Handler {
	return &OrdersHandler{repository}
}

func (h *OrdersHandler) RegisterRoutes(router *echo.Group, _ map[string]echo.MiddlewareFunc) {
	router.GET(constants.OrdersPath, h.List)
	router.POST(constants.OrdersPath, h.Create)
	router.GET(constants.OrderWithIDPath, h.Retrieve)
	router.PATCH(constants.OrderWithIDPath, h.PartialUpdate)
	router.DELETE(constants.OrderWithIDPath, h.Destroy)
}

func (h *OrdersHandler) List(c echo.Context) error {
	orders, err := h.repository.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"orders": orders}

	return c.JSON(http.StatusOK, response)
}

func (h *OrdersHandler) Create(c echo.Context) error {
	var order entities.Order
	if err := c.Bind(&order); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.repository.Create(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, order)
}

func (h *OrdersHandler) Retrieve(c echo.Context) error {
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

func (h *OrdersHandler) PartialUpdate(c echo.Context) error {
	var id uint64

	err := echo.PathParamsBinder(c).Uint64("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	var order entities.Order

	if err := c.Bind(&order); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	err = h.repository.PartialUpdate(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *OrdersHandler) Destroy(c echo.Context) error {
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
