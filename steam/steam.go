// Package steam provides a scraper for Steam Store featured deals.
//
// It fetches the current featured categories from Steam, identifies free games
// (100% discount) and significant deals (>50% discount), and returns structured data.
package steam

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"free-games-tracker/httpclient"
	"free-games-tracker/model"
)

const (
	// apiURL is the Steam featured categories endpoint.
	// cc=ID sets the country to Indonesia for IDR pricing.
	apiURL = "https://store.steampowered.com/api/featuredcategories/?cc=ID&l=english"

	// storeBaseURL is the base URL for Steam app pages.
	storeBaseURL = "https://store.steampowered.com/app/"

	// minDiscountPercent is the minimum discount to qualify as a "deal".
	minDiscountPercent = 50
)

// ── API Response Types ───────────────────────────────────────────────────────

type apiResponse struct {
	Specials *apiCategory `json:"specials"`
	TopSellers *apiCategory `json:"top_sellers"`
}

type apiCategory struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Items []apiItem `json:"items"`
}

type apiItem struct {
	ID              int    `json:"id"`
	Type            int    `json:"type"`
	Name            string `json:"name"`
	Discounted      bool   `json:"discounted"`
	DiscountPercent int    `json:"discount_percent"`
	OriginalPrice   *int   `json:"original_price"`
	FinalPrice      *int   `json:"final_price"`
	Currency        string `json:"currency"`
	LargeCapsuleImage string `json:"large_capsule_image"`
	HeaderImage     string `json:"header_image"`
	SmallCapsuleImage string `json:"small_capsule_image"`
	DiscountExpiration *int64 `json:"discount_expiration"`
}

// ── Public API ───────────────────────────────────────────────────────────────

// ScrapeResult holds both free games and deals found on Steam.
type ScrapeResult struct {
	FreeGames []model.Game
	Deals     []model.Deal
}

// FetchDeals fetches featured specials from Steam and returns:
//   - Free games (100% discount)
//   - Deals with >50% discount
func FetchDeals(ctx context.Context) (*ScrapeResult, error) {
	items, err := fetchSpecials(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching steam specials: %w", err)
	}

	now := time.Now().UTC()
	scrapedAt := now.Format(time.RFC3339)

	result := &ScrapeResult{}

	for i := range items {
		item := &items[i]

		if !item.Discounted || item.DiscountPercent == 0 {
			continue
		}

		// Skip expired deals.
		if item.DiscountExpiration != nil {
			expiry := time.Unix(*item.DiscountExpiration, 0)
			if expiry.Before(now) {
				continue
			}
		}

		thumbnail := pickThumbnail(item)
		url := buildStoreURL(item)

		// 100% discount = free game
		if item.DiscountPercent == 100 {
			result.FreeGames = append(result.FreeGames, model.Game{
				Title:         item.Name,
				Platform:      "Steam",
				OriginalPrice: formatSteamPrice(item.OriginalPrice, item.Currency),
				Thumbnail:     thumbnail,
				URL:           url,
				StartDate:     "",
				EndDate:       formatExpiration(item.DiscountExpiration),
				ScrapedAt:     scrapedAt,
			})
			continue
		}

		// >50% discount = deal
		if item.DiscountPercent >= minDiscountPercent {
			result.Deals = append(result.Deals, model.Deal{
				Title:           item.Name,
				Platform:        "Steam",
				OriginalPrice:   formatSteamPrice(item.OriginalPrice, item.Currency),
				DiscountedPrice: formatSteamPrice(item.FinalPrice, item.Currency),
				DiscountPercent: item.DiscountPercent,
				Thumbnail:       thumbnail,
				URL:             url,
				ScrapedAt:       scrapedAt,
			})
		}
	}

	return result, nil
}

// ── HTTP Fetch ───────────────────────────────────────────────────────────────

// fetchSpecials performs the HTTP request with retry and returns the specials items.
func fetchSpecials(ctx context.Context) ([]apiItem, error) {
	resp, err := httpclient.Do(ctx, "GET", apiURL)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	if apiResp.Specials == nil {
		return nil, nil
	}

	return apiResp.Specials.Items, nil
}

// ── Helpers ──────────────────────────────────────────────────────────────────

// formatSteamPrice converts a Steam price integer (in cents) to a readable string.
// Steam prices are in the smallest currency unit (e.g., 13799900 = IDR 137,999).
func formatSteamPrice(price *int, currency string) string {
	if price == nil {
		return "N/A"
	}

	p := *price
	if p == 0 {
		return "Free"
	}

	// Steam IDR prices are in full rupiah (no decimal).
	// Other currencies may use cents (e.g., USD 2999 = $29.99).
	switch currency {
	case "IDR":
		// IDR prices from Steam are already in rupiah units (e.g., 13799900 for IDR 137,999)
		// But sometimes they come as smaller units. Handle both.
		if p >= 100 {
			p = p / 100 // Convert from Steam's internal format
		}
		return fmt.Sprintf("IDR %s", formatWithThousands(p))
	default:
		// For other currencies, divide by 100 for decimal.
		whole := p / 100
		frac := p % 100
		if frac == 0 {
			return fmt.Sprintf("%s %s", currency, formatWithThousands(whole))
		}
		return fmt.Sprintf("%s %s.%02d", currency, formatWithThousands(whole), frac)
	}
}

// pickThumbnail selects the best available thumbnail.
func pickThumbnail(item *apiItem) string {
	if item.HeaderImage != "" {
		return item.HeaderImage
	}
	if item.LargeCapsuleImage != "" {
		return item.LargeCapsuleImage
	}
	if item.SmallCapsuleImage != "" {
		return item.SmallCapsuleImage
	}
	return ""
}

// buildStoreURL constructs the Steam store page URL.
func buildStoreURL(item *apiItem) string {
	if item.ID == 0 {
		return ""
	}
	return fmt.Sprintf("%s%d/", storeBaseURL, item.ID)
}

// formatExpiration converts a Unix timestamp to RFC3339 string.
func formatExpiration(ts *int64) string {
	if ts == nil {
		return ""
	}
	return time.Unix(*ts, 0).UTC().Format(time.RFC3339)
}

// formatWithThousands formats an integer with comma thousands separators.
func formatWithThousands(n int) string {
	if n < 0 {
		return "-" + formatWithThousands(-n)
	}

	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}

	var result []byte
	remainder := len(s) % 3
	if remainder > 0 {
		result = append(result, s[:remainder]...)
	}

	for i := remainder; i < len(s); i += 3 {
		if len(result) > 0 {
			result = append(result, ',')
		}
		result = append(result, s[i:i+3]...)
	}

	return string(result)
}
