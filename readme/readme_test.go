package readme

import (
	"strings"
	"testing"

	"free-games-tracker/model"
)

func TestGenerate_WithGamesDealsAndStats(t *testing.T) {
	games := []model.Game{
		{
			Title:         "Citizen Sleeper",
			Platform:      "Epic Games",
			OriginalPrice: "IDR 137,999",
			URL:           "https://store.epicgames.com/en-US/p/citizen-sleeper-944858",
			EndDate:       "2026-06-25T15:00:00.000Z",
		},
	}

	deals := []model.Deal{
		{
			Title:           "Hades",
			Platform:        "Steam",
			OriginalPrice:   "IDR 150,000",
			DiscountedPrice: "IDR 60,000",
			DiscountPercent: 60,
			URL:             "https://store.steampowered.com/app/1145360/",
		},
	}

	stats := Stats{
		TotalGamesTracked: 42,
		TotalValueSaved:   5200000,
		PlatformCounts:    map[string]int{"Epic Games": 30, "Steam": 12},
	}

	result := Generate(games, deals, stats, "2026-06-20 06:14")

	checks := []string{
		"# 🎮 Free Games Tracker",
		"**42** games tracked",
		"IDR 5,200,000",
		"Citizen Sleeper",
		"Epic Games",
		"Hades",
		"**-60%**",
		"~~IDR 150,000~~",
		"**IDR 60,000**",
	}

	for _, want := range checks {
		if !strings.Contains(result, want) {
			t.Errorf("README missing expected content: %q", want)
		}
	}
}

func TestGenerate_Empty(t *testing.T) {
	stats := Stats{PlatformCounts: map[string]int{}}
	result := Generate(nil, nil, stats, "2026-06-20 06:14")

	if !strings.Contains(result, "No free games found") {
		t.Error("should contain 'No free games found' message")
	}
	if !strings.Contains(result, "No big deals found") {
		t.Error("should contain 'No big deals found' message")
	}
}

func TestCalcStats(t *testing.T) {
	history := []model.Game{
		{Title: "Game A", Platform: "Epic Games", OriginalPrice: "IDR 137,999"},
		{Title: "Game B", Platform: "Epic Games", OriginalPrice: "IDR 100,000"},
		{Title: "Game C", Platform: "Steam", OriginalPrice: "IDR 50,000"},
	}

	stats := CalcStats(history)

	if stats.TotalGamesTracked != 3 {
		t.Errorf("TotalGamesTracked = %d, want 3", stats.TotalGamesTracked)
	}
	if stats.PlatformCounts["Epic Games"] != 2 {
		t.Errorf("Epic Games count = %d, want 2", stats.PlatformCounts["Epic Games"])
	}
	if stats.PlatformCounts["Steam"] != 1 {
		t.Errorf("Steam count = %d, want 1", stats.PlatformCounts["Steam"])
	}
	if stats.TotalValueSaved != 287999 {
		t.Errorf("TotalValueSaved = %d, want 287999", stats.TotalValueSaved)
	}
}

func TestParseIDRPrice(t *testing.T) {
	tests := []struct {
		name  string
		price string
		want  int64
	}{
		{"normal", "IDR 137,999", 137999},
		{"no_comma", "IDR 50000", 50000},
		{"free", "Free", 0},
		{"na", "N/A", 0},
		{"zero", "0", 0},
		{"empty", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseIDRPrice(tt.price)
			if got != tt.want {
				t.Errorf("parseIDRPrice(%q) = %d, want %d", tt.price, got, tt.want)
			}
		})
	}
}

func TestFormatDateDisplay(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", "N/A"},
		{"rfc3339", "2026-06-25T15:00:00Z", "Jun 25, 2026 15:00 UTC"},
		{"with_millis", "2026-06-25T15:00:00.000Z", "Jun 25, 2026 15:00 UTC"},
		{"invalid", "not-a-date", "not-a-date"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatDateDisplay(tt.input)
			if got != tt.want {
				t.Errorf("formatDateDisplay(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
