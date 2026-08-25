package api

import (
	"encoding/json"
	"go-url-shortener/database/db"
	"go-url-shortener/helpers"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Handler struct {
	queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{queries: queries}
}

func (h *Handler) CreateURL(w http.ResponseWriter, r *http.Request) {
	var req CreateUrlRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.DestinationURL == "" {
		http.Error(w, "your must provide a destination url", http.StatusBadRequest)
		return
	}

	var shortCode string

	for range generateCodeRetries {
		code, err := helpers.GenerateRandomCode()
		if err != nil {
			http.Error(w, "failed to generate short code", http.StatusInternalServerError)
			return
		}

		count, err := h.queries.CountURLByShortCode(r.Context(), code)
		if err != nil {
			http.Error(w, "failed to check short code", http.StatusInternalServerError)
			return
		}

		if count == 0 {
			shortCode = code
			break
		}
	}

	if shortCode == "" {
		http.Error(w, "failed to generate unique short code", http.StatusInternalServerError)
		return
	}

	expiresAt := time.Now().Add(24 * time.Hour)

	url, err := h.queries.CreateURL(r.Context(),
		db.CreateURLParams{
			ShortCode:      shortCode,
			DestinationUrl: req.DestinationURL,
			ExpiresAt: pgtype.Timestamptz{
				Time:  expiresAt,
				Valid: true,
			}})

	if err != nil {
		http.Error(w, "failed to create url", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(url)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) RedirectToURL(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	url, err := h.queries.GetURLByShortCode(r.Context(), code)

	if err != nil {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}

	if url.ExpiresAt.Valid && !time.Now().Before(url.ExpiresAt.Time) {
		http.Error(w, "url has expired", http.StatusGone)
		return
	}

	if err := h.queries.IncrementClickCount(r.Context(), code); err != nil {
		log.Println("failed to increment click count:", err)
	}

	http.Redirect(w, r, url.DestinationUrl, http.StatusFound)
}

func (h *Handler) GetURLStatistics(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	urlInformation, err := h.queries.GetURLInformation(r.Context(), code)

	if err != nil {
		http.Error(w, "url information not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(urlInformation)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}
