package epic

import (
	"testing"
	"time"
)

func TestFormatWithThousands(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"zero", 0, "0"},
		{"small", 42, "42"},
		{"hundreds", 999, "999"},
		{"thousands", 1000, "1,000"},
		{"typical_price", 137999, "137,999"},
		{"million", 1000000, "1,000,000"},
		{"large", 1234567890, "1,234,567,890"},
		{"negative", -5000, "-5,000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatWithThousands(tt.n)
			if got != tt.want {
				t.Errorf("formatWithThousands(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestExtractPrice(t *testing.T) {
	tests := []struct {
		name string
		elem apiElement
		want string
	}{
		{
			name: "nil_price",
			elem: apiElement{},
			want: "N/A",
		},
		{
			name: "formatted_price",
			elem: apiElement{
				Price: &apiPrice{
					TotalPrice: &apiTotalPrice{
						OriginalPrice: 137999,
						FmtPrice: &apiFmtPrice{
							OriginalPrice: "IDR\u00a0137,999.00",
						},
					},
				},
			},
			want: "IDR 137,999",
		},
		{
			name: "zero_price",
			elem: apiElement{
				Price: &apiPrice{
					TotalPrice: &apiTotalPrice{
						OriginalPrice: 0,
					},
				},
			},
			want: "0",
		},
		{
			name: "raw_integer_price",
			elem: apiElement{
				Price: &apiPrice{
					TotalPrice: &apiTotalPrice{
						OriginalPrice: 269000,
					},
				},
			},
			want: "IDR 269,000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractPrice(&tt.elem)
			if got != tt.want {
				t.Errorf("extractPrice() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExtractThumbnail(t *testing.T) {
	tests := []struct {
		name string
		elem apiElement
		want string
	}{
		{
			name: "no_images",
			elem: apiElement{},
			want: "",
		},
		{
			name: "prefers_thumbnail_type",
			elem: apiElement{
				KeyImages: []apiKeyImage{
					{Type: "OfferImageWide", URL: "https://wide.jpg"},
					{Type: "Thumbnail", URL: "https://thumb.jpg"},
				},
			},
			want: "https://thumb.jpg",
		},
		{
			name: "falls_back_to_first",
			elem: apiElement{
				KeyImages: []apiKeyImage{
					{Type: "Unknown", URL: "https://unknown.jpg"},
				},
			},
			want: "https://unknown.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractThumbnail(&tt.elem)
			if got != tt.want {
				t.Errorf("extractThumbnail() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildStoreURL(t *testing.T) {
	tests := []struct {
		name string
		elem apiElement
		want string
	}{
		{
			name: "from_offer_mappings",
			elem: apiElement{
				OfferMappings: []apiMapping{{PageSlug: "citizen-sleeper-944858"}},
			},
			want: "https://store.epicgames.com/en-US/p/citizen-sleeper-944858",
		},
		{
			name: "from_catalog_ns",
			elem: apiElement{
				CatalogNs: &apiCatalogNs{
					Mappings: []apiMapping{{PageSlug: "robobeat-5f084b"}},
				},
			},
			want: "https://store.epicgames.com/en-US/p/robobeat-5f084b",
		},
		{
			name: "no_mappings",
			elem: apiElement{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildStoreURL(&tt.elem)
			if got != tt.want {
				t.Errorf("buildStoreURL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFindActiveFreePromotion(t *testing.T) {
	now := time.Date(2026, 6, 20, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name   string
		elem   apiElement
		wantOK bool
	}{
		{
			name:   "no_promotions",
			elem:   apiElement{},
			wantOK: false,
		},
		{
			name: "active_free_promo",
			elem: apiElement{
				Promotions: &apiPromotions{
					PromotionalOffers: []apiPromoGroup{{
						PromotionalOffers: []apiOffer{{
							StartDate:       "2026-06-18T15:00:00.000Z",
							EndDate:         "2026-06-25T15:00:00.000Z",
							DiscountSetting: &apiDiscountSetting{DiscountPercentage: 0},
						}},
					}},
				},
			},
			wantOK: true,
		},
		{
			name: "expired_promo",
			elem: apiElement{
				Promotions: &apiPromotions{
					PromotionalOffers: []apiPromoGroup{{
						PromotionalOffers: []apiOffer{{
							StartDate:       "2026-06-01T15:00:00.000Z",
							EndDate:         "2026-06-10T15:00:00.000Z",
							DiscountSetting: &apiDiscountSetting{DiscountPercentage: 0},
						}},
					}},
				},
			},
			wantOK: false,
		},
		{
			name: "non_free_discount",
			elem: apiElement{
				Promotions: &apiPromotions{
					PromotionalOffers: []apiPromoGroup{{
						PromotionalOffers: []apiOffer{{
							StartDate:       "2026-06-18T15:00:00.000Z",
							EndDate:         "2026-06-25T15:00:00.000Z",
							DiscountSetting: &apiDiscountSetting{DiscountPercentage: 50},
						}},
					}},
				},
			},
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := findActiveFreePromotion(&tt.elem, now)
			if ok != tt.wantOK {
				t.Errorf("findActiveFreePromotion() ok = %v, want %v", ok, tt.wantOK)
			}
		})
	}
}
