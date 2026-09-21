package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func (u *User) GetUserOutput() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate, u.CreatedAt)
}

func NewUser(firstName, lastName, birthday string) (*User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthday,
		CreatedAt: time.Now(),
	}, nil
}
