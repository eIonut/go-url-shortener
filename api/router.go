package api

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /urls", handler.CreateURL)
	mux.HandleFunc("GET /{code}", handler.RedirectToURL)

	return mux
}
