package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "20")
}
