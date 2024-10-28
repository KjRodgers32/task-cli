package data

import (
	"fmt"
	"os"
	"encoding/csv"
	"time"
	"math"
)

type Data struct {
	ID string
	Task string
	Created time.Time
	Done bool
}

func (d *Data) DaysBetweenTasks() int {
	timeInBetween := time.Since(d.Created)
	if timeInBetween.Hours() <= 24.0 {
		return 0
	} else {
		return int(math.Floor(timeInBetween.Hours() / 24))
	}
}
// func createDataStore() {
// 	file, err := os.Create("task.csv")
// 	if err != nil {
// 		fmt.Printf("this is the error boy: %v", err)
// 		return
// 	}
// 	defer file.Close()
// }

func GetTasks() ([]Data, error) {
	file, err := os.Open("task.csv")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	defer file.Close()
	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error reading headers; %v", err)
	}

	var tasks []Data

	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			fmt.Println("Error reading row:", err)
			break
		}

		var task Data
		task.ID = row[0]
		task.Task = row[1]
		fmt.Println(row[2])
		date, err := time.Parse("2006-01-02",row[2])
		if err != nil {
			return nil, err
		}
		task.Created = date
		fmt.Sscan(row[3], &task.Done)

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func CreateTasks() {
	return 
}

func DeleteTasks() {
	return
}
