package models

type Token struct {
	UserID uint   `json:"userID"`
	Token  string `json:"token"`
	Phone  string `json:"phone"`
}
