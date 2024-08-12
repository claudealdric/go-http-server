package main

import (
	"log"
	"net/http"
)

func main() {
	store := &InMemoryPlayerStore{}
	server := &PlayerServer{store}
	log.Fatal(http.ListenAndServe(":8080", server))
}

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return 0
}

func (i *InMemoryPlayerStore) RecordWin(player string) {
}
