package poker

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()
	store, err := NewFileSystemStore(database)
	assertNoError(t, err)
	server := NewPlayerServer(store)
	player := "Pepper"
	wins := 3

	for range wins {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), strconv.Itoa(wins))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		assertLeague(
			t,
			getLeagueFromResponse(t, response.Body),
			[]Player{{"Pepper", wins}},
		)
	})
}
