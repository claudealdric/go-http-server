package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	player := strings.TrimPrefix(url, "/players/")
	score := p.store.GetPlayerScore(player)

	rw.WriteHeader(http.StatusNotFound)

	fmt.Fprint(rw, score)
}
