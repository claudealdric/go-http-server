package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/claudealdric/go-http-server"
)

const dbFileName = "game.db.json"
const port = 8080

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, err := poker.NewFileSystemStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
