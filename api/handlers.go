package api

import (
	"go-url-shortener/database/db"
	"net/http"
)

type Handler struct {
	queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{queries: queries}
}

func (h *Handler) CreateURL(w http.ResponseWriter, r *http.Request) {}
