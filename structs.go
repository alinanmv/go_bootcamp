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

func main() {
	userFirstName := getUserInput("first name:")
	userLastName := getUserInput("last name:")
	userBirthday := getUserInput("birthday:")

	var appUser User

	appUser = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthday,
		createdAt: time.Now(),
	}
	getUserOutput(appUser)
	saveUserToFile(appUser.firstName, appUser.lastName, appUser.birthDate, appUser.createdAt)
}

func getUserOutput(u User) {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func getUserInput(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
