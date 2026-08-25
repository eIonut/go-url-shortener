package api

type CreateUrlRequest struct {
	DestinationURL string `json:"destination_url"`
}

type CreateUrlResponse struct {
	ShortURL string `json:"shortURL"`
}
