package steam

import (
	"testing"
)

func TestFormatSteamPrice(t *testing.T) {
	tests := []struct {
		name     string
		price    *int
		currency string
		want     string
	}{
		{"nil_price", nil, "IDR", "N/A"},
		{"zero_price", intPtr(0), "IDR", "Free"},
		{"idr_price", intPtr(13799900), "IDR", "IDR 137,999"},
		{"idr_small", intPtr(10899900), "IDR", "IDR 108,999"},
		{"usd_price", intPtr(2999), "USD", "USD 29.99"},
		{"usd_whole", intPtr(6000), "USD", "USD 60"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatSteamPrice(tt.price, tt.currency)
			if got != tt.want {
				t.Errorf("formatSteamPrice(%v, %q) = %q, want %q", tt.price, tt.currency, got, tt.want)
			}
		})
	}
}

func TestPickThumbnail(t *testing.T) {
	tests := []struct {
		name string
		item apiItem
		want string
	}{
		{
			name: "prefers_header",
			item: apiItem{
				HeaderImage:       "https://header.jpg",
				LargeCapsuleImage: "https://large.jpg",
			},
			want: "https://header.jpg",
		},
		{
			name: "falls_back_to_large",
			item: apiItem{
				LargeCapsuleImage: "https://large.jpg",
			},
			want: "https://large.jpg",
		},
		{
			name: "falls_back_to_small",
			item: apiItem{
				SmallCapsuleImage: "https://small.jpg",
			},
			want: "https://small.jpg",
		},
		{
			name: "no_images",
			item: apiItem{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pickThumbnail(&tt.item)
			if got != tt.want {
				t.Errorf("pickThumbnail() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildStoreURL(t *testing.T) {
	tests := []struct {
		name string
		item apiItem
		want string
	}{
		{"valid_id", apiItem{ID: 730}, "https://store.steampowered.com/app/730/"},
		{"zero_id", apiItem{ID: 0}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildStoreURL(&tt.item)
			if got != tt.want {
				t.Errorf("buildStoreURL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatExpiration(t *testing.T) {
	ts := int64(1750867200)
	got := formatExpiration(&ts)
	if got == "" {
		t.Error("expected non-empty string")
	}

	got = formatExpiration(nil)
	if got != "" {
		t.Errorf("expected empty string for nil, got %q", got)
	}
}

func intPtr(n int) *int {
	return &n
}
