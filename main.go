package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Email    string
	Password string
}

// storage layer
var userStorage []User

func main() {
	fmt.Println("welcome to my app")

	//Flag
	command := flag.String("command", "no command", "help : command to run")
	flag.Parse()

	for {
		//Run
		runCommand(*command)

		//Run again
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}

	fmt.Println("User Storage : ", userStorage)
}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "user-register":
		userRegister()
	case "user-login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func createTask() {
	//Scanner, input
	scanner := bufio.NewScanner(os.Stdin)
	var name, category, dueDate string

	fmt.Println("please enter the task title")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the task Category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("please enter the task Due Date")
	scanner.Scan()
	dueDate = scanner.Text()

	//show
	fmt.Println("task : ", name, category, dueDate)
}

func createCategory() {
	//Scanner, input
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the color")
	scanner.Scan()
	color = scanner.Text()

	//show
	fmt.Println("category : ", title, color)
}

func userRegister() {
	//Scanner, input
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()
	//show without id
	fmt.Println("category : ", email, password)

	//unique
	id = email
	//show with id
	fmt.Println("User : ", id, email, password)
}

func login() {
	//Scanner, input
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter password")
	scanner.Scan()
	password = scanner.Text()

	//show without id
	fmt.Println("User : ", id, email, password)

	// Add to storage
	user := User{ID: len(userStorage) + 1, Email: email, Password: password}
	userStorage = append(userStorage, user)
}
