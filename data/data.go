package data

import (
	"fmt"
	"os"
)

type Data struct {
	ID string
	Task string
	Created string
	Done bool
}

func createDataStore() {
	file, err := os.Create("task.csv")
	if err != nil {
		fmt.Printf("this is the error boy: %v", err)
		return
	}
	defer file.Close()
}

func getTasks() ([]Data, error){
	file, err := os.Open("task.csv")
	fmt.Println(file)
	if err != nil {
		return nil, fmt.Errorf("error opening your task datastore: %v", err)
	}

	defer file.Close()

	// data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading your tasks: %v", err)
	}


	return nil, nil
}