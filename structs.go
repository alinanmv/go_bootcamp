package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const file = "users.txt"

func saveUserToFile(firstName, lastName, birthDate string, createdAt time.Time) {
	userText := fmt.Sprintf("Name: %v, last name: %v, birthday: %v, created at: %v\n",
		firstName, lastName, birthDate, createdAt)
	os.WriteFile(file, []byte(userText), 0644)
}

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) getUserOutput() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func newUser(firstName, lastName, birthday string) (*User, error) {
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
func main() {
	userFirstName := getUserInput("first name:")
	userLastName := getUserInput("last name:")
	userBirthday := getUserInput("birthday:")

	var appUser *User

	appUser, error := newUser(userFirstName, userLastName, userBirthday)
	if error != nil {
		fmt.Println(error)
		return
	}
	appUser.getUserOutput()
	saveUserToFile(appUser.firstName, appUser.lastName, appUser.birthDate, appUser.createdAt)
}

func getUserInput(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
