package emojistore

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"strings"

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

// Search finds emojis whose names contain the query string (case-insensitive)
func (s *Store) Search(query string) []Emoji {
	if query == "" {
		return []Emoji{}
	}

	query = strings.ToLower(query)
	emojis := make([]Emoji, 0)

	for _, e := range s.Emojis {
		if strings.Contains(strings.ToLower(e.Name), query) {
			emojis = append(emojis, e)
		}
	}

	return emojis
}

// GetSimilar finds emojis with names similar to the given name
// Returns emojis that share words or have partial matches
func (s *Store) GetSimilar(name string) []Emoji {
	if name == "" {
		return []Emoji{}
	}

	name = strings.ToLower(name)
	emojis := make([]Emoji, 0)
	words := strings.Fields(name)

	for _, e := range s.Emojis {
		emojiName := strings.ToLower(e.Name)

		// Skip exact matches
		if emojiName == name {
			continue
		}

		// Check if any word from the input name appears in the emoji name
		for _, word := range words {
			if len(word) > 2 && strings.Contains(emojiName, word) {
				emojis = append(emojis, e)
				break
			}
		}
	}

	return emojis
}
