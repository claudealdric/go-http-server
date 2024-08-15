package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.handleLeague))
	router.Handle("/players/", http.HandlerFunc(p.handlePlayers))

	p.Handler = router

	return p
}
func (p *PlayerServer) getLeagueTable() League {
	return []Player{
		{"Chris", 20},
	}
}

func (p *PlayerServer) handleLeague(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("content-type", jsonContentType)
	json.NewEncoder(rw).Encode(p.store.GetLeague())
}

func (p *PlayerServer) handlePlayers(rw http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.String(), "/players/")

	switch req.Method {
	case http.MethodPost:
		p.processWin(rw, player)
	case http.MethodGet:
		p.showScore(rw, player)
	}
}

func (p *PlayerServer) processWin(rw http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	rw.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(rw http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		rw.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(rw, score)
}
