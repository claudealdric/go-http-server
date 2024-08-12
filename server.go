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
	fmt.Fprint(rw, p.store.GetPlayerScore(player))
}
