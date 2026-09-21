package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}
type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "admin",
			lastName:  "admin",
			birthDate: time.Now().Format("2006-01-02"),
			createdAt: time.Now(),
		},
	}
}
func (u *User) GetUserOutput() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func New(firstName, lastName, birthday string) (*User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthday,
		createdAt: time.Now(),
	}, nil
}
