package data

import (
	"fmt"
	"os"
)

func createDataStore() {
	file, err := os.Open("task.csv")
	if err != nil {
		fmt.Printf("this is the error boy: %v", err)
		return
	}
	defer file.Close()
}