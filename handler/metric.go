package handler

import (
	"net/http"
	"strconv"

	"github.com/checkaayush/authware/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ListMetrics(c echo.Context) error {
	ctx := c.Request().Context()
	metrics, err := h.db.ListMetrics(ctx)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, metrics)
}

func (h *Handler) CreateMetric(c echo.Context) error {
	ctx := c.Request().Context()
	m := &model.Metric{}
	if err := c.Bind(m); err != nil {
		return err
	}

	metric, err := h.db.CreateMetric(ctx, m)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, metric)
}

func (h *Handler) DeleteMetric(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusBadRequest, nil)
	}

	err = h.db.DeleteMetric(ctx, id)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.NoContent(http.StatusNoContent)
}
