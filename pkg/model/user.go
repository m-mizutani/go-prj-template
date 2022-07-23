package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/m-mizutani/goerr"
)

type (
	UserID     string
	PersonName string
)

func NewUserID() UserID {
	return UserID(uuid.NewString())
}

type User struct {
	ID        UserID
	Name      PersonName
	CreatedAt time.Time
}

func NewUser(name PersonName) *User {
	return &User{
		ID:   NewUserID(),
		Name: name,
	}
}

func (x *User) Validate() error {
	if x.ID == "" {
		return goerr.Wrap(ErrInvalidInput, "User.ID is required")
	}

	return nil
}
