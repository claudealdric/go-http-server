package poker

import (
	"encoding/json"
	"io"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner. got %q, want %q", store.winCalls[0], winner)
	}
}

func AssertContentType(t *testing.T, got string, wantedContentType string) {
	t.Helper()
	if got != wantedContentType {
		t.Errorf(
			"response did not have content-type of application/json, got %v",
			got,
		)
	}
}

func AssertLeague(t *testing.T, got []Player, wantedLeague []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague) {
		t.Errorf("got %v, want %v", got, wantedLeague)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf(
			"Unable to parse response from server %q into slice of Player, '%v'",
			body,
			err,
		)
	}

	return
}
