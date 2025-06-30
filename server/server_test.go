package server

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cheatsnake/emojihub/emojistore"
	"github.com/julienschmidt/httprouter"
)

func TestEmojis(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.Emojis(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		expected := store.GetAll()
		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody []emojistore.Emoji
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		if len(expected) != len(jsonBody) {
			t.Errorf("expected body with %d emojis, but got %d", len(expected), len(jsonBody))
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.Emojis(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}

func TestEmojisByCategory(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		randCategory := store.Categories[rand.Intn(len(store.Categories))]
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.EmojisByCategory(rr, req, httprouter.Params{{Key: "category", Value: strings.ReplaceAll(randCategory, " ", "-")}})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		expected := store.GetAllByCategory(randCategory)
		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody []emojistore.Emoji
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		if len(expected) != len(jsonBody) {
			t.Errorf("expected body with %d emojis, but got %d", len(expected), len(jsonBody))
		}
	})

	t.Run("request to a non-existent category", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.EmojisByCategory(rr, req, httprouter.Params{{Key: "category", Value: "wtf"}})

		if rr.Result().StatusCode != http.StatusNotFound {
			t.Errorf("expected status 404, but got %d", rr.Result().StatusCode)
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.EmojisByCategory(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}

func TestEmojisByGroup(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		randGroup := store.Groups[rand.Intn(len(store.Groups))]
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.EmojisByGroup(rr, req, httprouter.Params{{Key: "group", Value: strings.ReplaceAll(randGroup, " ", "-")}})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		expected := store.GetAllByGroup(randGroup)
		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody []emojistore.Emoji
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		if len(expected) != len(jsonBody) {
			t.Errorf("expected body with %d emojis, but got %d", len(expected), len(jsonBody))
		}
	})

	t.Run("request to a non-existent group", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.EmojisByGroup(rr, req, httprouter.Params{{Key: "group", Value: "wtf"}})

		if rr.Result().StatusCode != http.StatusNotFound {
			t.Errorf("expected status 404, but got %d", rr.Result().StatusCode)
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.EmojisByGroup(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}

func TestRandomEmoji(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.RandomEmoji(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody emojistore.Emoji
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		if !(len(jsonBody.Name) > 0) {
			t.Errorf("expected emoji with defined name")
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.RandomEmoji(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}

func TestRandomEmojiByCategory(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		randCategory := store.Categories[rand.Intn(len(store.Categories))]
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.RandomEmojiByCategory(rr, req, httprouter.Params{{Key: "category", Value: strings.ReplaceAll(randCategory, " ", "-")}})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody emojistore.Emoji
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		if !(len(jsonBody.Name) > 0) {
			t.Errorf("expected emoji with defined name")
		}
	})

	t.Run("request to a non-existent category", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.RandomEmojiByCategory(rr, req, httprouter.Params{{Key: "category", Value: "wtf"}})

		if rr.Result().StatusCode != http.StatusNotFound {
			t.Errorf("expected status 404, but got %d", rr.Result().StatusCode)
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.RandomEmojiByCategory(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}

func TestRandomEmojiByGroup(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		randGroup := store.Groups[rand.Intn(len(store.Groups))]
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.RandomEmojiByGroup(rr, req, httprouter.Params{{Key: "group", Value: strings.ReplaceAll(randGroup, " ", "-")}})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody emojistore.Emoji
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		if !(len(jsonBody.Name) > 0) {
			t.Errorf("expected emoji with defined name")
		}
	})

	t.Run("request to a non-existent group", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.RandomEmojiByGroup(rr, req, httprouter.Params{{Key: "group", Value: "wtf"}})

		if rr.Result().StatusCode != http.StatusNotFound {
			t.Errorf("expected status 404, but got %d", rr.Result().StatusCode)
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.RandomEmojiByGroup(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}

func TestCategories(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.Categories(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		expected := store.Categories
		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody []string
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		if len(expected) != len(jsonBody) {
			t.Errorf("expected body with %d categories, but got %d", len(expected), len(jsonBody))
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.Categories(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}

func TestGroups(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.Groups(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		expected := store.Groups
		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody []string
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		if len(expected) != len(jsonBody) {
			t.Errorf("expected body with %d groups, but got %d", len(expected), len(jsonBody))
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.Groups(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}

func TestSearch(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid search request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/api/search?q=smile", nil)
		if err != nil {
			t.Error(err)
		}

		server.Search(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody []emojistore.Emoji
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		// Check that results contain "smile" in name
		if len(jsonBody) > 0 {
			found := false
			for _, emoji := range jsonBody {
				if strings.Contains(strings.ToLower(emoji.Name), "smile") {
					found = true
					break
				}
			}
			if !found {
				t.Error("expected at least one emoji with 'smile' in name")
			}
		}
	})

	t.Run("search with empty query", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/api/search", nil)
		if err != nil {
			t.Error(err)
		}

		server.Search(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, but got %d", rr.Result().StatusCode)
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/api/search?q=test", nil)
		if err != nil {
			t.Error(err)
		}

		server.Search(rr, req, httprouter.Params{})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}

func TestSimilar(t *testing.T) {
	store := emojistore.New()
	server := New(store)

	t.Run("valid similar request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.Similar(rr, req, httprouter.Params{{Key: "name", Value: "smiling face"}})

		if rr.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status 200, but got %d", rr.Result().StatusCode)
		}

		defer rr.Result().Body.Close()

		body, err := io.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}

		var jsonBody []emojistore.Emoji
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			t.Error(err)
		}

		// Results should be an array (could be empty)
		if jsonBody == nil {
			t.Error("expected valid JSON array response")
		}
	})

	t.Run("request with not allowed method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "", nil)
		if err != nil {
			t.Error(err)
		}

		server.Similar(rr, req, httprouter.Params{{Key: "name", Value: "test"}})

		if rr.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, but got %d", rr.Result().StatusCode)
		}
	})
}
