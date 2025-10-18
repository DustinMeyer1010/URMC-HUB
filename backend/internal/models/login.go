package models

// Login information for the user
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
