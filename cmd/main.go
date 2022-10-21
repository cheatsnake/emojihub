package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/cheatsnake/emojihub/emojistore"
	"github.com/cheatsnake/emojihub/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	server := server.Server{Store: emojistore.New()}

	http.HandleFunc("/api/all/", server.Emojis)
	http.HandleFunc("/api/random/", server.RandomEmoji)

	log.Fatal(http.ListenAndServe(":5005", nil))
}
