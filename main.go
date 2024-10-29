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
				tasks, err := data.GetTasks()
				if err != nil {
					fmt.Printf("there was an error retrieving your tasks: %v\n", err)
					return
				}

				fmt.Println("ID			Task			Created			Done")

				for _, task := range tasks {
					daysPassed, err := task.DaysBetweenTasks()
					if err != nil {
						fmt.Printf("there was an error calculating days between tasks: %v", err)
					}
					fmt.Printf("%s %s %d days ago %v\n", task.ID, task.Task, daysPassed, task.Done)
				}

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