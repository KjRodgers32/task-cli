package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("What would you like to do...")
		fmt.Println("1.	Read list of tasks")
		fmt.Println("2.	Create a task")
		fmt.Println("3.	Delete as task")
		fmt.Println("4.	Quit")

		var response string
		fmt.Scanln(&response)

		switch response {
			case "1":
				data.getTasks()
			case "2":
				data.createTask()
			case "3":
				data.deleteTask()
			case "4":
				return
			default:
				fmt.Println("Please select a vaild response...")
		}
	}
}