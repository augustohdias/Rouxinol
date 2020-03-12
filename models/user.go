package models

// User type
type User struct {
	UserID       string   `json:"user_id,omitempty"`
	Username     string   `json:"username,omitempty"`
	Alias        string   `json:"alias,omitempty"`
	SessionToken string   `json:"session_token,omitempty"`
	Following    []string `json:"following,omitempty"`
}
