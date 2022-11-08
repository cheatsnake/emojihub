package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) Emojis(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	emojis := s.Store.GetAll()
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) EmojisByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	category := strings.Replace(ps.ByName("category"), "-", " ", -1)
	emojis := s.Store.GetAllByCategory(category)
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) EmojisByGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	group := strings.Replace(ps.ByName("group"), "-", " ", -1)
	emojis := s.Store.GetAllByGroup(group)
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) RandomEmoji(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	emoji := s.Store.GetRandom()
	json, _ := json.Marshal(emoji)

	w.Write(json)
}

func (s *Server) RandomEmojiByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	category := strings.Replace(ps.ByName("category"), "-", " ", -1)
	emoji := s.Store.GetRandomByCategory(category)
	json, _ := json.Marshal(emoji)

	w.Write(json)
}

func (s *Server) RandomEmojiByGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	group := strings.Replace(ps.ByName("group"), "-", " ", -1)
	emoji := s.Store.GetRandomByGroup(group)
	json, _ := json.Marshal(emoji)

	w.Write(json)
}