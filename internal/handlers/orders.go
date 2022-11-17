package handlers

import (
	"interview/domain/constants"
	"interview/domain/entities"
	"interview/internal/repositories"
	"interview/internal/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type OrdersHandler struct {
	repository repositories.OrdersRepository
}

func NewOrdersHandler(repository repositories.OrdersRepository) Handler {
	return &OrdersHandler{repository}
}

func (h *OrdersHandler) RegisterRoutes(router *echo.Group, mws map[string]echo.MiddlewareFunc) {
	userRoutes := router.Group("", mws[constants.User])
	userRoutes.GET(constants.OrdersPath, h.UserList)
	userRoutes.POST(constants.OrdersPath, h.UserCreate)
	userRoutes.GET(constants.OrderWithIDPath, h.UserRetrieve)
	userRoutes.POST(constants.OrderWithIDPath+"/cancel", h.CancelOrder)

	internalRoutes := router.Group("/internal", mws[constants.Admin])
	internalRoutes.GET(constants.OrdersPath, h.List)
	internalRoutes.POST(constants.OrdersPath, h.Create)
	internalRoutes.GET(constants.OrderWithIDPath, h.Retrieve)
	internalRoutes.PUT(constants.OrderWithIDPath, h.Update)
	internalRoutes.DELETE(constants.OrderWithIDPath, h.Destroy)
}

func (h *OrdersHandler) List(c echo.Context) error {
	order := &entities.Order{}
	orders, err := h.repository.List(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"orders": orders}

	return c.JSON(http.StatusOK, response)
}

func (h *OrdersHandler) Create(c echo.Context) error {
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

func (h *OrdersHandler) Retrieve(c echo.Context) error {
	var id uint

	err := echo.PathParamsBinder(c).Uint("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order := &entities.Order{}
	order.ID = id

	err = h.repository.Retrieve(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func (h *OrdersHandler) Update(c echo.Context) error {
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

func (h *OrdersHandler) UserList(c echo.Context) error {
	order := &entities.Order{}
	order.User = GetUserID(c)

	orders, err := h.repository.List(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"orders": orders}

	return c.JSON(http.StatusOK, response)
}

func (h *OrdersHandler) UserCreate(c echo.Context) error {
	order := &entities.Order{}
	if err := c.Bind(order); err != nil {
		return err
	}

	order.User = GetUserID(c)
	order.Status = constants.DefaultOrderStatus
	order.Rembolso = false

	err := h.repository.Create(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, order)
}

func (h *OrdersHandler) UserRetrieve(c echo.Context) error {
	var id uint

	err := echo.PathParamsBinder(c).Uint("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order := &entities.Order{}
	order.User = GetUserID(c)
	order.ID = id

	c.Logger().Print("user ", order.User)
	c.Logger().Debug("user ", order.User)

	err = h.repository.Retrieve(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func (h *OrdersHandler) CancelOrder(c echo.Context) error {
	var id uint

	err := echo.PathParamsBinder(c).Uint("id", &id).BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	order := &entities.Order{}
	order.User = GetUserID(c)
	order.ID = id

	err = h.repository.Retrieve(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if utils.Contains(order.Status, constants.InvalidCancelStatus) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid status")
	}

	order.Status = constants.CancelOrderStatus

	t := time.Now()
	if elapsed := t.Sub(order.CreatedAt); elapsed <= time.Duration(120)*time.Second {
		order.Refund = true
	}

	err = h.repository.Update(c.Request().Context(), order)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
