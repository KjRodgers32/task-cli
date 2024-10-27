package main

import (
	"fmt"
	"github.com/KjRodgers32/task-cli/data"
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
				data.GetTasks()
			case "2":
				data.CreateTasks()
			case "3":
				data.DeleteTasks()
			case "4":
				return
			default:
				fmt.Println("Please select a vaild response...")
		}
	}
}