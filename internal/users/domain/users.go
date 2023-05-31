package users

import "time"

type UserID int64

type User struct {
	ID        UserID
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
}
