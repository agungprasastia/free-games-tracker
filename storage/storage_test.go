package storage

import (
	"os"
	"path/filepath"
	"testing"

	"free-games-tracker/model"
)

func TestLoadGames_NonExistent(t *testing.T) {
	games, err := LoadGames("/nonexistent/path/games.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(games) != 0 {
		t.Errorf("expected empty slice, got %d games", len(games))
	}
}

func TestLoadGames_EmptyFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty.json")
	if err := os.WriteFile(path, []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}

	games, err := LoadGames(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(games) != 0 {
		t.Errorf("expected empty slice, got %d games", len(games))
	}
}

func TestSaveAndLoadGames(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sub", "games.json")

	want := []model.Game{
		{
			Title:         "Test Game",
			Platform:      "Epic Games",
			OriginalPrice: "IDR 100,000",
		},
	}

	if err := SaveGames(path, want); err != nil {
		t.Fatalf("SaveGames: %v", err)
	}

	got, err := LoadGames(path)
	if err != nil {
		t.Fatalf("LoadGames: %v", err)
	}

	if len(got) != 1 {
		t.Fatalf("expected 1 game, got %d", len(got))
	}
	if got[0].Title != "Test Game" {
		t.Errorf("Title = %q, want %q", got[0].Title, "Test Game")
	}
}

func TestMergeHistory_NoDuplicates(t *testing.T) {
	current := []model.Game{
		{Title: "Game A", Platform: "Epic Games"},
	}
	history := []model.Game{
		{Title: "Game A", Platform: "Epic Games"},
	}

	added := MergeHistory(current, &history)

	if added != 0 {
		t.Errorf("expected 0 new entries, got %d", added)
	}
	if len(history) != 1 {
		t.Errorf("expected 1 total, got %d", len(history))
	}
}

func TestMergeHistory_AddsNew(t *testing.T) {
	current := []model.Game{
		{Title: "Game B", Platform: "Epic Games"},
	}
	history := []model.Game{
		{Title: "Game A", Platform: "Epic Games"},
	}

	added := MergeHistory(current, &history)

	if added != 1 {
		t.Errorf("expected 1 new entry, got %d", added)
	}
	if len(history) != 2 {
		t.Errorf("expected 2 total, got %d", len(history))
	}
	if history[1].Title != "Game B" {
		t.Errorf("history[1].Title = %q, want %q", history[1].Title, "Game B")
	}
}

func TestMergeHistory_SameTitleDifferentPlatform(t *testing.T) {
	current := []model.Game{
		{Title: "Game A", Platform: "Steam"},
	}
	history := []model.Game{
		{Title: "Game A", Platform: "Epic Games"},
	}

	added := MergeHistory(current, &history)

	if added != 1 {
		t.Errorf("expected 1 new entry (different platform), got %d", added)
	}
	if len(history) != 2 {
		t.Errorf("expected 2 total, got %d", len(history))
	}
}
