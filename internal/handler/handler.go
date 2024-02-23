package handler

import (
	"net/http"
	"news_feed/internal/repo"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo *repo.Repo
}

func NewHandler(r *repo.Repo) Handler {
	return Handler{
		repo: r,
	}
}

func (h *Handler) GetItems(c echo.Context) error {
	items, err := h.repo.ReadItems()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, items)
}
