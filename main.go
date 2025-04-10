package main

import (
	"fmt"
	"github.com/KjRodgers32/task-cli/data"
	"log"
	"strconv"
	"strings"
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
				log.Fatal("error reading tasks", err)
			}
			maxRowLength, err := data.MaxColLength(tasks)
			if err != nil {
				log.Fatal("error getting max row length")
			}
			spaceBuffer := strings.Repeat(" ", maxRowLength)
			fmt.Printf("ID%sTask%sCreated%sDone\n", spaceBuffer, spaceBuffer, spaceBuffer)
			for _, task := range tasks {
				taskSpaceBuffers := taskSpaceBuffer(task, maxRowLength)
				fmt.Printf("%s%s%s%s%v%s%v\n", task.ID, strings.Repeat(" ", taskSpaceBuffers.IdLength), task.Task, strings.Repeat(" ", taskSpaceBuffers.TaskLength), task.Created.Format("2006-01-02"), strings.Repeat(" ", taskSpaceBuffers.CreaedLength), task.Done)
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

func taskSpaceBuffer(task data.Data, maxRowLength int) data.RowLength {
	createdDataString := task.Created.String()
	doneString := strconv.FormatBool(task.Done)
	return data.RowLength{
		IdLength:     ifNeg(maxRowLength - len(task.ID)),
		TaskLength:   ifNeg(maxRowLength - len(task.Task)),
		CreaedLength: ifNeg(maxRowLength - len(createdDataString)),
		DoneLength:   ifNeg(maxRowLength - len(doneString)),
	}
}

func ifNeg(nbr int) int {
	if nbr < 0 {
		return 0
	}
	return nbr
}
