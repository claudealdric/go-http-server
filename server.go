package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		player := strings.TrimPrefix(req.URL.String(), "/players/")

		switch req.Method {
		case http.MethodPost:
			p.processWin(rw, player)
		case http.MethodGet:
			p.showScore(rw, player)
		}
	}))

	router.ServeHTTP(rw, req)
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
