package models

// Post type
type Post struct {
	ID         string `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	Alias      string `json:"alias,omitempty"`
	Text       string `json:"text,omitempty"`
	Attachment string `json:"attachment,omitempty"`
}
