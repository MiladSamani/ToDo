package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	email    string
	password string
}

var userStorage []User

func main() {
	fmt.Println("hello, Welcome to my ToDO Application!")
	command := flag.String("command", "no command", "command to run")
	flag.Parse()
	runCommand(*command)
	fmt.Printf("User Storage %v\n : ", userStorage)
	for {
		fmt.Println("try another command")
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTasks()
	case "create-category":
		createCategory()
	case "registered-user":
		registeredUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func createTasks() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, DueDate, category string
	fmt.Println("please Enter your Tasks :")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("please Enter your category :")
	scanner.Scan()
	category = scanner.Text()
	fmt.Println("please Enter your Due date :")
	scanner.Scan()
	DueDate = scanner.Text()
	fmt.Println("tasks : ", name, DueDate, category)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string
	fmt.Println("please Enter your category title :")
	scanner.Scan()
	title = scanner.Text()
	fmt.Println("please Enter your category color :")
	scanner.Scan()
	color = scanner.Text()
	fmt.Println("category : ", title, color)
}

func registeredUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var ID, Password, email string
	fmt.Println("please Enter your email :")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("please Enter your password :")
	scanner.Scan()
	Password = scanner.Text()
	fmt.Println(Password, email)
	ID = email
	fmt.Println(ID)
	user := User{
		ID: len(userStorage) + 1, email: email, password: Password}
	userStorage = append(userStorage, user)
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	flag.Parse()
	var ID, Password, email string
	fmt.Println("please Enter your email :")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("please Enter your password :")
	scanner.Scan()
	Password = scanner.Text()
	fmt.Println(email, Password)
	fmt.Println(ID, email, Password)
}
