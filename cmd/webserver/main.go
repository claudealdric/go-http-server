package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	poker "github.com/claudealdric/go-http-server"
)

const dbFileName = "game.db.json"
const port = 8080

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, err := poker.NewFileSystemStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), server); err != nil {
		log.Fatalf("could not listen on port %d %v", port, err)
	}
}
