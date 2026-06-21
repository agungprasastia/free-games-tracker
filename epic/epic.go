// Package epic provides a scraper for the Epic Games Store free games promotions API.
//
// It fetches the current free game listings, parses promotional offers to identify
// games with a 100% discount, and returns structured Game data.
package epic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"free-games-tracker/httpclient"
	"free-games-tracker/model"
)

const (
	// apiURL is the Epic Games Store free games promotions endpoint.
	apiURL = "https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=ID&allowCountries=ID"

	// storeBaseURL is the base URL for Epic Games Store product pages.
	storeBaseURL = "https://store.epicgames.com/en-US/p/"
)

// ── API Response Types ───────────────────────────────────────────────────────
// These structs model the Epic Games Store API JSON response. Fields that are
// not needed are intentionally omitted to keep the deserialization lean.

type apiResponse struct {
	Data *apiData `json:"data"`
}

type apiData struct {
	Catalog *apiCatalog `json:"Catalog"`
}

type apiCatalog struct {
	SearchStore *apiSearchStore `json:"searchStore"`
}

type apiSearchStore struct {
	Elements []apiElement `json:"elements"`
}

type apiElement struct {
	Title         string          `json:"title"`
	KeyImages     []apiKeyImage   `json:"keyImages"`
	CatalogNs     *apiCatalogNs   `json:"catalogNs"`
	OfferMappings []apiMapping    `json:"offerMappings"`
	Price         *apiPrice       `json:"price"`
	Promotions    *apiPromotions  `json:"promotions"`
}

type apiKeyImage struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type apiCatalogNs struct {
	Mappings []apiMapping `json:"mappings"`
}

type apiMapping struct {
	PageSlug string `json:"pageSlug"`
}

type apiPrice struct {
	TotalPrice *apiTotalPrice `json:"totalPrice"`
}

type apiTotalPrice struct {
	OriginalPrice int          `json:"originalPrice"`
	FmtPrice      *apiFmtPrice `json:"fmtPrice"`
}

type apiFmtPrice struct {
	OriginalPrice string `json:"originalPrice"`
}

type apiPromotions struct {
	PromotionalOffers []apiPromoGroup `json:"promotionalOffers"`
}

type apiPromoGroup struct {
	PromotionalOffers []apiOffer `json:"promotionalOffers"`
}

type apiOffer struct {
	StartDate       string           `json:"startDate"`
	EndDate         string           `json:"endDate"`
	DiscountSetting *apiDiscountSetting `json:"discountSetting"`
}

type apiDiscountSetting struct {
	DiscountPercentage int `json:"discountPercentage"`
}

// ── activePromo holds the dates of a verified free promotion ─────────────────

type activePromo struct {
	StartDate string
	EndDate   string
}

// ── Public API ───────────────────────────────────────────────────────────────

// FetchFreeGames fetches and returns all currently free games from the
// Epic Games Store. It uses the provided context for cancellation and timeout.
func FetchFreeGames(ctx context.Context) ([]model.Game, error) {
	elements, err := fetchElements(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching epic games api: %w", err)
	}

	now := time.Now().UTC()
	scrapedAt := now.Format(time.RFC3339)
	var games []model.Game

	for i := range elements {
		elem := &elements[i]

		promo, ok := findActiveFreePromotion(elem, now)
		if !ok {
			continue
		}

		games = append(games, model.Game{
			Title:         elem.Title,
			Platform:      "Epic Games",
			OriginalPrice: extractPrice(elem),
			Thumbnail:     extractThumbnail(elem),
			URL:           buildStoreURL(elem),
			StartDate:     promo.StartDate,
			EndDate:       promo.EndDate,
			ScrapedAt:     scrapedAt,
		})
	}

	return games, nil
}

// ── HTTP Fetch ───────────────────────────────────────────────────────────────

// fetchElements performs the HTTP request with retry and returns the raw API elements.
func fetchElements(ctx context.Context) ([]apiElement, error) {
	resp, err := httpclient.Do(ctx, "GET", apiURL)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	if apiResp.Data == nil || apiResp.Data.Catalog == nil || apiResp.Data.Catalog.SearchStore == nil {
		return nil, nil
	}

	return apiResp.Data.Catalog.SearchStore.Elements, nil
}

// ── Promotion Detection ─────────────────────────────────────────────────────

// findActiveFreePromotion checks whether the given element has an active
// promotion with a 100% discount (i.e., completely free). It returns the
// promotion dates and true if found, or zero-value and false otherwise.
func findActiveFreePromotion(elem *apiElement, now time.Time) (activePromo, bool) {
	if elem.Promotions == nil {
		return activePromo{}, false
	}

	for _, group := range elem.Promotions.PromotionalOffers {
		for _, offer := range group.PromotionalOffers {
			// Only consider 100% discount offers.
			if offer.DiscountSetting == nil || offer.DiscountSetting.DiscountPercentage != 0 {
				continue
			}

			// Skip expired promotions.
			if offer.EndDate != "" {
				endTime, err := time.Parse(time.RFC3339, offer.EndDate)
				if err == nil && endTime.Before(now) {
					continue
				}
			}

			return activePromo{
				StartDate: offer.StartDate,
				EndDate:   offer.EndDate,
			}, true
		}
	}

	return activePromo{}, false
}

// ── Data Extraction Helpers ──────────────────────────────────────────────────

// extractPrice returns the human-readable original price of the game.
// It tries the formatted price first, then falls back to the raw integer.
func extractPrice(elem *apiElement) string {
	if elem.Price == nil || elem.Price.TotalPrice == nil {
		return "N/A"
	}

	tp := elem.Price.TotalPrice

	// Prefer the pre-formatted price string from the API.
	if tp.FmtPrice != nil && tp.FmtPrice.OriginalPrice != "" {
		price := tp.FmtPrice.OriginalPrice
		price = strings.ReplaceAll(price, "\u00a0", " ") // non-breaking space
		price = strings.ReplaceAll(price, ".00", "")
		return price
	}

	// Fall back to raw integer price.
	if tp.OriginalPrice == 0 {
		return "0"
	}
	return fmt.Sprintf("IDR %s", formatWithThousands(tp.OriginalPrice))
}

// extractThumbnail selects the best thumbnail URL from the element's key images.
// It prioritizes image types in a specific order for consistent display.
func extractThumbnail(elem *apiElement) string {
	if len(elem.KeyImages) == 0 {
		return ""
	}

	// Priority order for image types.
	preferredTypes := []string{
		"Thumbnail",
		"OfferImageWide",
		"DieselStoreFrontWide",
		"OfferImageTall",
	}

	for _, ptype := range preferredTypes {
		for _, img := range elem.KeyImages {
			if img.Type == ptype && img.URL != "" {
				return img.URL
			}
		}
	}

	// Fall back to the first available image.
	for _, img := range elem.KeyImages {
		if img.URL != "" {
			return img.URL
		}
	}

	return ""
}

// buildStoreURL constructs the Epic Games Store product page URL.
// It checks offerMappings first (newer API format), then catalogNs mappings.
func buildStoreURL(elem *apiElement) string {
	// Try offerMappings first (newer API format).
	for _, m := range elem.OfferMappings {
		if m.PageSlug != "" {
			return storeBaseURL + m.PageSlug
		}
	}

	// Try catalogNs mappings.
	if elem.CatalogNs != nil {
		for _, m := range elem.CatalogNs.Mappings {
			if m.PageSlug != "" {
				return storeBaseURL + m.PageSlug
			}
		}
	}

	return ""
}

// ── Formatting Utilities ────────────────────────────────────────────────────

// formatWithThousands formats an integer with comma thousands separators.
// Example: 137999 → "137,999"
func formatWithThousands(n int) string {
	if n < 0 {
		return "-" + formatWithThousands(-n)
	}

	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}

	var result strings.Builder
	remainder := len(s) % 3
	if remainder > 0 {
		result.WriteString(s[:remainder])
	}

	for i := remainder; i < len(s); i += 3 {
		if result.Len() > 0 {
			result.WriteByte(',')
		}
		result.WriteString(s[i : i+3])
	}

	return result.String()
}
