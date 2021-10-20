package models

type URL struct {
	URL string `json:"url"   binding:"required"`
}

type ShortURLResponce struct {
	ShortURL string `json:"short_url"`
}

type FullURLResponce struct {
	FullURL string `json:"full_url"`
}
