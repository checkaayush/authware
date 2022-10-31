package handler

import (
	"net/http"
	"strconv"

	"github.com/checkaayush/authware/model"
	"github.com/labstack/echo/v4"
)

// ListApps lists apps on the platform
func (h *Handler) ListApps(c echo.Context) error {
	ctx := c.Request().Context()
	apps, err := h.db.ListApps(ctx)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, apps)
}

// CreateApp creates a new app on the platform
func (h *Handler) CreateApp(c echo.Context) error {
	ctx := c.Request().Context()
	a := &model.App{}
	if err := c.Bind(a); err != nil {
		return err
	}

	app, err := h.db.CreateApp(ctx, a)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, app)
}

// DeleteApp deletes an app from the platform
func (h *Handler) DeleteApp(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusBadRequest, nil)
	}

	err = h.db.DeleteApp(ctx, id)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.NoContent(http.StatusNoContent)
}
