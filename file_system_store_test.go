package main

import (
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
		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

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
		store, err := NewFileSystemStore(database)

		assertNoError(t, err)
		assertScoreEquals(t, store.GetPlayerScore("Chris"), 33)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		store, err := NewFileSystemStore(database)
		assertNoError(t, err)
		store.RecordWin("Chris")
		assertScoreEquals(t, store.GetPlayerScore("Chris"), 34)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		store, err := NewFileSystemStore(database)
		assertNoError(t, err)
		player := "Pepper"
		store.RecordWin(player)

		assertScoreEquals(t, store.GetPlayerScore(player), 1)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemStore(database)
		assertNoError(t, err)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
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
