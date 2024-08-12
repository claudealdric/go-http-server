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
	switch req.Method {
	case http.MethodPost:
		p.processWin(rw)
	case http.MethodGet:
		p.showScore(rw, req)
	}
}

func (p *PlayerServer) processWin(rw http.ResponseWriter) {
	p.store.RecordWin("Bob")
	rw.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(rw http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.String(), "/players/")
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		rw.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(rw, score)
}
