package user

import "time"

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	Role           string
	Token          string
	AvatarFileName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
