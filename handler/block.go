package handler

import (
	"net/http"
	"strconv"

	"github.com/checkaayush/authware/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ListBlocks(c echo.Context) error {
	ctx := c.Request().Context()
	blocks, err := h.db.ListBlocks(ctx)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, blocks)
}

func (h *Handler) CreateBlock(c echo.Context) error {
	ctx := c.Request().Context()
	b := &model.Block{}
	if err := c.Bind(b); err != nil {
		return err
	}

	block, err := h.db.CreateBlock(ctx, b)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, block)
}

func (h *Handler) DeleteBlock(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusBadRequest, nil)
	}

	err = h.db.DeleteBlock(ctx, id)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.NoContent(http.StatusNoContent)
}
