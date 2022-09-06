package forms

import "time"

type AccessToken struct {
	ID        uint
	Hash      string
	TypeID    uint
	UserID    uint
	ExpiresAt time.Time
}
