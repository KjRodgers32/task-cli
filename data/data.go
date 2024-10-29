package data

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func Max(nums ...int) int {
	largest := int(math.Inf(-1))
	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}
	return largest
}

func maxColLength(tasks []Data) int {
	var maxLength int
	for _, task := range tasks {
		maxLength = Max(maxLength, len(task.ID), len(task.Task), len(task.Created.Format("2006-01-02")), len(strconv.FormatBool(task.Done)))
	}
	return maxLength
}

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
		date, err := time.Parse("2006-01-02", row[2])
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
