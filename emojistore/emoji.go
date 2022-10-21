package emojistore

type Emoji struct {
	Name     string   `json:"name"`
	Category string   `json:"category"`
	Group    string   `json:"group"`
	HtmlCode []string `json:"htmlCode"`
	Unicode  []string `json:"unicode"`
}
