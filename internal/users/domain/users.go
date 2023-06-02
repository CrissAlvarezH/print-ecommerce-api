package users

import "time"

type UserID int64

type User struct {
	ID        UserID
	Name      string
	Email     string
	Phone     string
	IsActive  bool
	CreatedAt time.Time
	Addresses []Address
	Scopes    []Scope
}

type ScopeName string

type Scope struct {
	Name ScopeName
}

type AddressID int64

type Address struct {
	ID            AddressID
	Department    string
	City          string
	Address       string
	ReceiverPhone string
	ReceiverName  string
}
