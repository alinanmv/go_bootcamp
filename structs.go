package main

import (
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
func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}
func newUser(firstName, lastName, birthday string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthday,
		createdAt: time.Now(),
	}
}
func main() {
	userFirstName := getUserInput("first name:")
	userLastName := getUserInput("last name:")
	userBirthday := getUserInput("birthday:")

	var appUser *User

	appUser = newUser(userFirstName, userLastName, userBirthday)
	appUser.getUserOutput()
	saveUserToFile(appUser.firstName, appUser.lastName, appUser.birthDate, appUser.createdAt)
}

func getUserInput(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
