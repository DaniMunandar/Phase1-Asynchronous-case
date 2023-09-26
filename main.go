package main

import (
	"fmt"
	"time"
)

type Login struct {
	User    string
	Message string
}

func sendNotif(user string, message string) {
	time.Sleep(2 * time.Second)
	fmt.Println(user, message)
}

func main() {
	loging := []Login{
		{User: "[INFO] - ", Message: "User 123 logged in"},
		{User: "[WARN] - ", Message: "Memory uase id high"},
		{User: "[ERROR] - ", Message: "Failed to fetch data from API"},
	}

	for _, v := range loging {
		go sendNotif(v.User, v.Message)
	}
	fmt.Printf("Application continues after logging...\n")
	time.Sleep(3 * time.Second)
	fmt.Println("main application finished")
}
