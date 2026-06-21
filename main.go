// Free Games Tracker — CLI scraper for free games from Epic Games Store & Steam.
//
// Usage:
//
//	free-games-tracker [flags]
//
// Flags:
//
//	-data-dir string     Path to the data directory (default "data")
//	-generate-readme     Generate/update README.md
//	-readme-path string  Path to the README file (default "README.md")
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"free-games-tracker/epic"
	"free-games-tracker/model"
	"free-games-tracker/readme"
	"free-games-tracker/steam"
	"free-games-tracker/storage"
)

// scrapeResult holds the output from a single platform scraper.
type scrapeResult struct {
	Platform string
	Games    []model.Game
	Deals    []model.Deal
	Err      error
	Duration time.Duration
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	dataDir := flag.String("data-dir", "data", "Path to the data directory")
	genReadme := flag.Bool("generate-readme", false, "Generate/update README.md")
	readmePath := flag.String("readme-path", "README.md", "Path to the README file")
	flag.Parse()

	gamesPath := filepath.Join(*dataDir, "games.json")
	dealsPath := filepath.Join(*dataDir, "deals.json")
	historyPath := filepath.Join(*dataDir, "history.json")

	printHeader()

	fmt.Println("🔍 Scraping platforms in parallel...")
	fmt.Println()

	ctx := context.Background()
	results := scrapeAll(ctx)

	var allFreeGames []model.Game
	var allDeals []model.Deal

	for _, r := range results {
		prefix := fmt.Sprintf("   [%s]", r.Platform)

		if r.Err != nil {
			fmt.Fprintf(os.Stderr, "%s ⚠  Error: %v\n", prefix, r.Err)
			continue
		}

		fmt.Printf("%s ✅ Done in %s\n", prefix, r.Duration.Round(time.Millisecond))

		if len(r.Games) > 0 {
			fmt.Printf("%s    🎮 %d free game(s):\n", prefix, len(r.Games))
			for _, g := range r.Games {
				fmt.Printf("%s       • %s (%s)\n", prefix, g.Title, g.OriginalPrice)
			}
			allFreeGames = append(allFreeGames, r.Games...)
		}

		if len(r.Deals) > 0 {
			fmt.Printf("%s    🏷️  %d deal(s) >50%% off:\n", prefix, len(r.Deals))
			for _, d := range r.Deals {
				fmt.Printf("%s       • %s (-%d%% → %s)\n", prefix, d.Title, d.DiscountPercent, d.DiscountedPrice)
			}
			allDeals = append(allDeals, r.Deals...)
		}

		if len(r.Games) == 0 && len(r.Deals) == 0 {
			fmt.Printf("%s    ℹ  Nothing found.\n", prefix)
		}

		fmt.Println()
	}

	fmt.Println("💾 Saving data...")

	if err := storage.SaveGames(gamesPath, allFreeGames); err != nil {
		return fmt.Errorf("saving games: %w", err)
	}
	fmt.Printf("   ✅ Saved %d free game(s) to %s\n", len(allFreeGames), gamesPath)

	if err := storage.SaveDeals(dealsPath, allDeals); err != nil {
		return fmt.Errorf("saving deals: %w", err)
	}
	fmt.Printf("   ✅ Saved %d deal(s) to %s\n", len(allDeals), dealsPath)

	history, err := storage.LoadGames(historyPath)
	if err != nil {
		return fmt.Errorf("loading history: %w", err)
	}

	newEntries := storage.MergeHistory(allFreeGames, &history)

	if err := storage.SaveGames(historyPath, history); err != nil {
		return fmt.Errorf("saving history: %w", err)
	}
	fmt.Printf("   ✅ History: %d total entries (%d new)\n", len(history), newEntries)

	if *genReadme {
		updatedAt := time.Now().UTC().Format("2006-01-02 15:04")
		stats := readme.CalcStats(history)
		content := readme.Generate(allFreeGames, allDeals, stats, updatedAt)

		if err := os.WriteFile(*readmePath, []byte(content), 0o644); err != nil {
			return fmt.Errorf("writing README: %w", err)
		}
		fmt.Printf("   ✅ Generated %s\n", *readmePath)
	}

	fmt.Println()
	fmt.Println("🎉 Done!")
	return nil
}

func scrapeAll(ctx context.Context) []scrapeResult {
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		results []scrapeResult
	)

	scrapers := []struct {
		name string
		fn   func(context.Context) (*scrapeResult, error)
	}{
		{"Epic Games", scrapeEpic},
		{"Steam", scrapeSteam},
	}

	wg.Add(len(scrapers))

	for _, s := range scrapers {
		go func(name string, fn func(context.Context) (*scrapeResult, error)) {
			defer wg.Done()

			start := time.Now()
			result, err := fn(ctx)

			if err != nil {
				mu.Lock()
				results = append(results, scrapeResult{
					Platform: name,
					Err:      err,
					Duration: time.Since(start),
				})
				mu.Unlock()
				return
			}

			result.Platform = name
			result.Duration = time.Since(start)

			mu.Lock()
			results = append(results, *result)
			mu.Unlock()
		}(s.name, s.fn)
	}

	wg.Wait()
	return results
}

func scrapeEpic(ctx context.Context) (*scrapeResult, error) {
	games, err := epic.FetchFreeGames(ctx)
	if err != nil {
		return nil, err
	}
	return &scrapeResult{Games: games}, nil
}

func scrapeSteam(ctx context.Context) (*scrapeResult, error) {
	sr, err := steam.FetchDeals(ctx)
	if err != nil {
		return nil, err
	}
	return &scrapeResult{
		Games: sr.FreeGames,
		Deals: sr.Deals,
	}, nil
}

func printHeader() {
	fmt.Println("🎮 Free Games Tracker — Go Edition v0.2.0")
	fmt.Println("═══════════════════════════════════════════")
	fmt.Println()
}
