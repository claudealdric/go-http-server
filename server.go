package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(rw http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	player := strings.TrimPrefix(url, "/players/")

	if player == "Pepper" {
		fmt.Fprint(rw, "20")
	}

	if player == "Floyd" {
		fmt.Fprint(rw, "10")
	}
}
