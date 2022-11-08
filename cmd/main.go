package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/cheatsnake/emojihub/emojistore"
	"github.com/cheatsnake/emojihub/server"
	"github.com/julienschmidt/httprouter"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	port := "5005"
	store := emojistore.New()
	server := server.New(store)
	router := httprouter.New()

	router.GET("/api/all", server.Emojis)
	router.GET("/api/all/category/:category", server.EmojisByCategory)
	router.GET("/api/all/group/:group", server.EmojisByGroup)

	router.GET("/api/random", server.RandomEmoji)
	router.GET("/api/random/category/:category", server.RandomEmojiByCategory)
	router.GET("/api/random/group/:group", server.RandomEmojiByGroup)

	fmt.Printf("Server starts on the port %s... \n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
