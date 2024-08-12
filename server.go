package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(rw http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	player := strings.TrimPrefix(url, "/players/")
	fmt.Fprint(rw, GetPlayerScore(player))

}

func GetPlayerScore(player string) int {
	if player == "Pepper" {
		return 20
	}

	if player == "Floyd" {
		return 10
	}

	return 0
}
