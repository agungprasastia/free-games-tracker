// Package storage provides JSON file persistence for game data.
//
// It handles loading, saving, and deduplicating game records stored as
// JSON arrays on disk.
package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"free-games-tracker/model"
)

// ── Games ────────────────────────────────────────────────────────────────────

// LoadGames reads a JSON file and returns the parsed slice of games.
// Returns an empty slice if the file does not exist or is empty.
func LoadGames(path string) ([]model.Game, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Game{}, nil
		}
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}

	if len(data) == 0 {
		return []model.Game{}, nil
	}

	var games []model.Game
	if err := json.Unmarshal(data, &games); err != nil {
		return nil, fmt.Errorf("parsing JSON from %s: %w", path, err)
	}

	return games, nil
}

// SaveGames writes the given games slice to a JSON file with pretty formatting.
// It creates the parent directory if it does not exist.
func SaveGames(path string, games []model.Game) error {
	return saveJSON(path, games)
}

// MergeHistory appends new games to the history, skipping duplicates.
// A game is considered duplicate if it has the same (title, platform) pair.
// Returns the number of new entries added.
func MergeHistory(current []model.Game, history *[]model.Game) int {
	type key struct {
		Title    string
		Platform string
	}

	existing := make(map[key]struct{}, len(*history))
	for _, g := range *history {
		existing[key{g.Title, g.Platform}] = struct{}{}
	}

	added := 0
	for _, g := range current {
		k := key{g.Title, g.Platform}
		if _, ok := existing[k]; ok {
			continue
		}
		*history = append(*history, g)
		existing[k] = struct{}{}
		added++
	}

	return added
}

// ── Deals ────────────────────────────────────────────────────────────────────

// LoadDeals reads a JSON file and returns the parsed slice of deals.
// Returns an empty slice if the file does not exist or is empty.
func LoadDeals(path string) ([]model.Deal, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Deal{}, nil
		}
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}

	if len(data) == 0 {
		return []model.Deal{}, nil
	}

	var deals []model.Deal
	if err := json.Unmarshal(data, &deals); err != nil {
		return nil, fmt.Errorf("parsing JSON from %s: %w", path, err)
	}

	return deals, nil
}

// SaveDeals writes the given deals slice to a JSON file with pretty formatting.
func SaveDeals(path string, deals []model.Deal) error {
	return saveJSON(path, deals)
}

// ── Internal Helpers ─────────────────────────────────────────────────────────

// saveJSON writes any serializable value to a JSON file.
func saveJSON(path string, v any) error {
	if err := ensureDir(path); err != nil {
		return err
	}

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling JSON: %w", err)
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}

	return nil
}

// ensureDir creates the parent directory for the given file path if needed.
func ensureDir(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("creating directory %s: %w", dir, err)
	}
	return nil
}
