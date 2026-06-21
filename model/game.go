// Package model defines the core data types used throughout the application.
package model

// Game represents a free game scraped from a game store.
type Game struct {
	Title         string `json:"title"`
	Platform      string `json:"platform"`
	OriginalPrice string `json:"original_price"`
	Thumbnail     string `json:"thumbnail"`
	URL           string `json:"url"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	ScrapedAt     string `json:"scraped_at"`
}

// Deal represents a game with a significant discount (>50%).
type Deal struct {
	Title           string `json:"title"`
	Platform        string `json:"platform"`
	OriginalPrice   string `json:"original_price"`
	DiscountedPrice string `json:"discounted_price"`
	DiscountPercent int    `json:"discount_percent"`
	Thumbnail       string `json:"thumbnail"`
	URL             string `json:"url"`
	ScrapedAt       string `json:"scraped_at"`
}
