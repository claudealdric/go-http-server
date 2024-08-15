package main

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`,
	)
	defer cleanDatabase()

	t.Run("league from a reader", func(t *testing.T) {
		store := NewFileSystemStore(database)

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		store := NewFileSystemStore(database)

		assertScoreEquals(t, store.GetPlayerScore("Chris"), 33)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		store := NewFileSystemStore(database)
		store.RecordWin("Chris")

		assertScoreEquals(t, store.GetPlayerScore("Chris"), 34)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		store := NewFileSystemStore(database)
		player := "Pepper"
		store.RecordWin(player)

		assertScoreEquals(t, store.GetPlayerScore(player), 1)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tempFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempFile.Write([]byte(initialData))

	removeFile := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}
