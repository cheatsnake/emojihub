package emojistore

import "testing"

func TestGetAll(t *testing.T) {
	store := New()
	emojis := store.GetAll()

	for _, emoji := range emojis {
		if len(emoji.Name) < 1 {
			t.Errorf("each emoji must have a name")
		}
		if len(emoji.Category) < 1 {
			t.Errorf("each emoji must have a category")
		}
		if len(emoji.Group) < 1 {
			t.Errorf("each emoji must have a group")
		}
		if len(emoji.HtmlCode) < 1 {
			t.Errorf("each emoji must have at least one html code")
		}
		if len(emoji.Unicode) < 1 {
			t.Errorf("each emoji must have at least one unicode")
		}
	}
}

func TestGetAllByCategory(t *testing.T) {
	store := New()

	for _, category := range store.Categories {
		emojis := store.GetAllByCategory(category)
		for _, emoji := range emojis {
			if emoji.Category != category {
				t.Errorf("emoji must have a category = %s, but got %s", category, emoji.Category)
			}
		}
	}
}

func TestGetAllByGroup(t *testing.T) {
	store := New()

	for _, group := range store.Groups {
		emojis := store.GetAllByGroup(group)
		for _, emoji := range emojis {
			if emoji.Group != group {
				t.Errorf("emoji must have a group = %s, but got %s", group, emoji.Group)
			}
		}
	}
}

func TestGetRandom(t *testing.T) {
	store := New()
	emoji := store.GetRandom()

	if len(emoji.Name) < 1 {
		t.Errorf("emoji must have a name")
	}
	if len(emoji.Category) < 1 {
		t.Errorf("emoji must have a category")
	}
	if len(emoji.Group) < 1 {
		t.Errorf("emoji must have a group")
	}
	if len(emoji.HtmlCode) < 1 {
		t.Errorf("emoji must have at least one html code")
	}
	if len(emoji.Unicode) < 1 {
		t.Errorf("emoji must have at least one unicode")
	}
}

func TestGetRandomByCategory(t *testing.T) {
	store := New()

	for _, category := range store.Categories {
		emoji := store.GetRandomByCategory(category)

		if emoji.Category != category {
			t.Errorf("emoji must have a category = %s, but got %s", category, emoji.Category)
		}
	}
}

func TestGetRandomByGroup(t *testing.T) {
	store := New()

	for _, group := range store.Groups {
		emoji := store.GetRandomByGroup(group)

		if emoji.Group != group {
			t.Errorf("emoji must have a group = %s, but got %s", group, emoji.Group)
		}
	}
}
