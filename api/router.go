package api

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /urls", handler.CreateURL)
	mux.HandleFunc("GET /{code}", handler.RedirectToURL)
	mux.HandleFunc("GET /urls/{code}", handler.GetURLStatistics)

	return mux
}
