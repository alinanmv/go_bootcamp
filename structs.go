package main

import (
	"fmt"
	"os"
	"time"

	"example/bootcamp/user"
)

const file = "users.txt"

func saveUserToFile(firstName, lastName, birthDate string, createdAt time.Time) {
	userText := fmt.Sprintf("Name: %v, last name: %v, birthday: %v, created at: %v\n",
		firstName, lastName, birthDate, createdAt)
	os.WriteFile(file, []byte(userText), 0644)
}

func main() {
	userFirstName := getUserInput("first name:")
	userLastName := getUserInput("last name:")
	userBirthday := getUserInput("birthday:")

	var appUser *user.User

	appUser, error := user.NewUser(userFirstName, userLastName, userBirthday)
	if error != nil {
		fmt.Println(error)
		return
	}
	appUser.GetUserOutput()
	saveUserToFile(appUser.FirstName, appUser.LastName, appUser.BirthDate, appUser.CreatedAt)
}

func getUserInput(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
