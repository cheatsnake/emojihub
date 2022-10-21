package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cheatsnake/emojihub/emojistore"
)

type Server struct {
	Store *emojistore.Store
}

func (s *Server) Emojis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	emojis := s.Store.GetAll()
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) EmojisByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category := strings.Split(r.URL.Path, "/")[3]
	emojis := s.Store.GetAllByCategory(category)
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) EmojisByCategoryGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category := strings.Split(r.URL.Path, "/")[3]
	group := strings.Split(r.URL.Path, "/")[4]
	emojis := s.Store.GetAllByCategoryGroup(category, group)
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) RandomEmoji(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	emoji := s.Store.GetRandom()
	json, _ := json.Marshal(emoji)

	w.Write(json)
}

func (s *Server) RandomEmojiByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category := strings.Split(r.URL.Path, "/")[3]
	emoji := s.Store.GetRandomByCategory(category)
	json, _ := json.Marshal(emoji)

	w.Write(json)
}

func (s *Server) RandomEmojiByCategoryGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category := strings.Split(r.URL.Path, "/")[3]
	group := strings.Split(r.URL.Path, "/")[4]
	emoji := s.Store.GetRandomByCategoryGroup(category, group)
	json, _ := json.Marshal(emoji)

	w.Write(json)
}