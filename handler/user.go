package handler

import (
	"net/http"
	"strconv"

	"github.com/checkaayush/authware/model"
	"github.com/labstack/echo/v4"
)

// ListUsers lists users on the platform
func (h *Handler) ListUsers(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := h.db.ListUsers(ctx)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, users)
}

// InviteUser invites a new user onto the platform
func (h *Handler) InviteUser(c echo.Context) error {
	ctx := c.Request().Context()
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	user, err := h.db.CreateUser(ctx, u)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, user)
}

// DeleteUser deletes an existing user from the platform
func (h *Handler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusBadRequest, nil)
	}

	err = h.db.DeleteUser(ctx, id)
	if err != nil {
		// TODO: Send appropriate error message
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.NoContent(http.StatusNoContent)
}
