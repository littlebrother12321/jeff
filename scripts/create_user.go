package main

import (
	"bufio"
	"fmt"
	"os"
	"firstbee/models"
	"firstbee/utils"
	"regexp"
	"strings"

	"golang.org/x/term"
)

func readName(reader *bufio.Reader) string {
	fmt.Print("Full Name: ")
	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return name
}

func readEmail(reader *bufio.Reader) string {
	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	// Define a regex pattern (e.g., to match an email address)
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile the regex
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(email) {
		fmt.Println("The input is NOT a valid email address.")
		os.Exit(1)
	}
	return email
}

func readPassword() string {
	const PASSWORD_LENGTH = 8
	fmt.Print("Password: ")
	passbyte, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	fmt.Print("\nConfirm Password: ")
	passbyte2, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	password, password2 := string(passbyte), string(passbyte2)
	if password != password2 {
		fmt.Println("Passwords do not match.")
		os.Exit(1)
	}
	if len(password) < PASSWORD_LENGTH {
		fmt.Printf("Password must be %d characters or more in length.\n", PASSWORD_LENGTH)
		os.Exit(1)
	}
	return password
}

func main() {
	models.InitDB()

	fmt.Println("Create User")
	reader := bufio.NewReader(os.Stdin)
	name := readName(reader)
	email := readEmail(reader)
	password := readPassword()

	// Create User
	user := models.User{
		Email:    email,
		Name:     name,
		Password: password,
	}
	_, err := utils.SaveUser(&user)
	if err != nil {
		fmt.Println("Error saving user: " + err.Error())
		os.Exit(1)
	} else {
		fmt.Println("User created successfully!")
	}
}
