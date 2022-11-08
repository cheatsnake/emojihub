package server

import (
	"github.com/cheatsnake/emojihub/emojistore"
)

type Server struct {
	Store *emojistore.Store
}

func New(es *emojistore.Store) *Server {
	return &Server{Store: es}
}
