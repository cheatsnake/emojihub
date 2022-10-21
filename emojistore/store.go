package emojistore

import (
	"encoding/json"
	"math/rand"
	"os"
)

type Store struct {
	Emojis     []Emoji
}

func New() *Store {
	newStore := Store{}
	file, err := os.Open("./data/emojibase.min.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Token()

	emoji := Emoji{}
	for decoder.More() {
		decoder.Decode(&emoji)
		newStore.Emojis = append(newStore.Emojis, emoji)
	}

	return &newStore
}

func (s *Store) GetAll() []Emoji {
	return s.Emojis
}

func (s *Store) GetAllByCategory(ctg string) []Emoji {
	emojis := make([]Emoji, 0, len(s.Emojis) / 2)

	for _, e := range s.Emojis {
		if e.Category == ctg {
			emojis = append(emojis, e)
		}
	}

	return emojis
}

func (s *Store) GetAllByCategoryGroup(ctg string, grp string) []Emoji {
	filtered := s.GetAllByCategory(ctg)
	emojis := make([]Emoji, 0, len(filtered))

	for _, e := range filtered {
		if e.Category == ctg {
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
	idx := rand.Intn(len(emojis))

	return emojis[idx]
}

func (s *Store) GetRandomByCategoryGroup(ctg string, grp string) Emoji {
	filtered := s.GetAllByCategory(ctg)
	emojis := make([]Emoji, 0, len(filtered))
	idx := rand.Intn(len(emojis))

	return emojis[idx]
}
