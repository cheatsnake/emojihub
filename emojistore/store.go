package emojistore

import (
	_ "embed"
	"encoding/json"
	"math/rand"

	"golang.org/x/exp/slices"
)

//go:embed data/emojibase.min.json
var emojibase []byte

type Emoji struct {
	Name     string   `json:"name"`
	Category string   `json:"category"`
	Group    string   `json:"group"`
	HtmlCode []string `json:"htmlCode"`
	Unicode  []string `json:"unicode"`
}

type Store struct {
	Emojis     []Emoji
	Categories []string
	Groups     []string
}

func New() *Store {
	var emojis []Emoji
	newStore := Store{}

	err := json.Unmarshal(emojibase, &emojis)
	if err != nil {
		panic(err)
	}

	for _, emoji := range emojis {
		if !slices.Contains(newStore.Categories, emoji.Category) {
			newStore.Categories = append(newStore.Categories, emoji.Category)
		}
		if !slices.Contains(newStore.Groups, emoji.Group) {
			newStore.Groups = append(newStore.Groups, emoji.Group)
		}
	}

	newStore.Emojis = emojis

	return &newStore
}

func (s *Store) GetAll() []Emoji {
	return s.Emojis
}

func (s *Store) GetAllByCategory(ctg string) []Emoji {
	emojis := make([]Emoji, 0, len(s.Emojis)/2)

	for _, e := range s.Emojis {
		if e.Category == ctg {
			emojis = append(emojis, e)
		}
	}

	return emojis
}

func (s *Store) GetAllByGroup(grp string) []Emoji {
	emojis := make([]Emoji, 0, len(s.Emojis)/2)

	for _, e := range s.Emojis {
		if e.Group == grp {
			emojis = append(emojis, e)
		}
	}

	return emojis
}

func (s *Store) GetRandom() Emoji {
	idx := rand.Intn(len(s.Emojis))
	return s.Emojis[idx]
}

func (s *Store) GetRandomByCategory(ctg string) Emoji {
	emojis := s.GetAllByCategory(ctg)

	if len(emojis) < 1 {
		return Emoji{}
	}

	idx := rand.Intn(len(emojis))

	return emojis[idx]
}

func (s *Store) GetRandomByGroup(grp string) Emoji {
	emojis := s.GetAllByGroup(grp)

	if len(emojis) < 1 {
		return Emoji{}
	}

	idx := rand.Intn(len(emojis))

	return emojis[idx]
}
