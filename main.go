package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID       int
	title    string
	dueDate  string
	category string
	isDone   bool
	userID   int
}

// User storage layer
var userStorage []User

// Task storage layer
var taskStorage []Task

var authenticatedUser *User

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
}

func runCommand(command string) {
	if command != "user-register" && command != "exit" && authenticatedUser == nil {
		login()
		if authenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "user-register":
		userRegister()
	case "user-login":
		login()
	case "task-list":
		listTask()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func createTask() {
	//Scanner, input
	scanner := bufio.NewScanner(os.Stdin)
	var title, category, dueDate string

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task Category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("please enter the task Due Date")
	scanner.Scan()
	dueDate = scanner.Text()

	task := Task{
		ID:       len(taskStorage) + 1,
		title:    title,
		dueDate:  dueDate,
		category: category,
		isDone:   false,
		userID:   authenticatedUser.ID,
	}
	taskStorage = append(taskStorage, task)
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
	var id, email, password, name string

	fmt.Println("please enter your name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	//unique
	id = email
	//show with id
	fmt.Println(name, id, email, password)

	// Add to storage
	user := User{ID: len(userStorage) + 1, Email: email, Password: password, Name: name}
	userStorage = append(userStorage, user)
}

func login() {
	fmt.Println("login process")
	//Scanner, input
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter password")
	scanner.Scan()
	password = scanner.Text()

	//get the email and password from the client
	//if there is a user record with corresponding data allow the user to continue
	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			authenticatedUser = &user
			break
		}
	}
	if authenticatedUser == nil {
		fmt.Println("the email or password is not correct")
	}
}

func (u User) print() {
	fmt.Println("User :", u.ID, u.Email, u.Name)
}

func listTask() {
	for _, task := range taskStorage {
		if task.userID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}
