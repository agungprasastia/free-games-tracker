// Package readme generates Markdown README files from game data.
package readme

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"free-games-tracker/model"
)

// Stats holds computed statistics about the game history.
type Stats struct {
	TotalGamesTracked int
	TotalValueSaved   int64 // in IDR
	PlatformCounts    map[string]int
}

// CalcStats computes statistics from the game history.
func CalcStats(history []model.Game) Stats {
	stats := Stats{
		TotalGamesTracked: len(history),
		PlatformCounts:    make(map[string]int),
	}

	for _, g := range history {
		stats.PlatformCounts[g.Platform]++
		stats.TotalValueSaved += parseIDRPrice(g.OriginalPrice)
	}

	return stats
}

// Generate creates a Markdown README string displaying free games, deals, and stats.
func Generate(games []model.Game, deals []model.Deal, stats Stats, updatedAt string) string {
	var b strings.Builder

	b.WriteString("# 🎮 Free Games Tracker\n\n")
	b.WriteString("Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.\n\n")
	b.WriteString(fmt.Sprintf("_Last updated: %s UTC_\n\n", updatedAt))

	// ── Stats Badge ──────────────────────────────────────────────────────
	if stats.TotalGamesTracked > 0 {
		b.WriteString(fmt.Sprintf("> 📊 **%d** games tracked", stats.TotalGamesTracked))
		if stats.TotalValueSaved > 0 {
			b.WriteString(fmt.Sprintf(" · **IDR %s** total value saved", formatWithThousands(stats.TotalValueSaved)))
		}
		for platform, count := range stats.PlatformCounts {
			b.WriteString(fmt.Sprintf(" · %s: %d", platform, count))
		}
		b.WriteString("\n\n")
	}

	// ── Free Games Section ───────────────────────────────────────────────
	b.WriteString("## 🔥 Current free games\n\n")

	if len(games) == 0 {
		b.WriteString("_No free games found right now. Check back later!_\n")
	} else {
		b.WriteString("| Game | Platform | Normal Price | Available Until | Link |\n")
		b.WriteString("|------|----------|-------------|----------------|------|\n")

		for _, g := range games {
			endDisplay := formatDateDisplay(g.EndDate)
			b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | [Claim](%s) |\n",
				g.Title, g.Platform, g.OriginalPrice, endDisplay, g.URL))
		}
	}

	// ── Deals Section ────────────────────────────────────────────────────
	b.WriteString("\n## 🏷️ Steam deals (>50% off)\n\n")

	if len(deals) == 0 {
		b.WriteString("_No big deals found right now._\n")
	} else {
		b.WriteString("| Game | Discount | Original | Sale Price | Link |\n")
		b.WriteString("|------|----------|----------|------------|------|\n")

		for _, d := range deals {
			b.WriteString(fmt.Sprintf("| %s | **-%d%%** | ~~%s~~ | **%s** | [View](%s) |\n",
				d.Title, d.DiscountPercent, d.OriginalPrice, d.DiscountedPrice, d.URL))
		}
	}

	// ── Footer ───────────────────────────────────────────────────────────
	b.WriteString("\n## 📦 Data\n\n")
	b.WriteString("- [`data/games.json`](data/games.json) — current free games\n")
	b.WriteString("- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)\n")
	b.WriteString("- [`data/history.json`](data/history.json) — all free games ever tracked\n")
	b.WriteString("\n## 🤖 How it works\n\n")
	b.WriteString("GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, ")
	b.WriteString("updates the data files, and commits the changes automatically.\n\n")
	b.WriteString("Built with **Go** 🐹 for simplicity and performance.\n")

	return b.String()
}

// ── Helpers ──────────────────────────────────────────────────────────────────

func formatDateDisplay(isoDate string) string {
	if isoDate == "" {
		return "N/A"
	}

	t, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05.000Z", isoDate)
		if err != nil {
			return isoDate
		}
	}

	return t.UTC().Format("Jan 02, 2006 15:04 UTC")
}

var priceRegex = regexp.MustCompile(`[\d,]+`)

func parseIDRPrice(price string) int64 {
	if price == "" || price == "0" || price == "Free" || price == "N/A" {
		return 0
	}

	match := priceRegex.FindString(price)
	if match == "" {
		return 0
	}

	cleaned := strings.ReplaceAll(match, ",", "")
	n, err := strconv.ParseInt(cleaned, 10, 64)
	if err != nil {
		return 0
	}

	return n
}

func formatWithThousands(n int64) string {
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
