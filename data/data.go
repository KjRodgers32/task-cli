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

func MaxColLength(tasks []Data) (int, error) {
	var maxLength int
	for _, task := range tasks {
		day, err := task.DaysBetweenTasks()
		if err != nil {
			return 0, err
		}
		maxLength = Max(maxLength, len(task.ID), len(task.Task), len(fmt.Sprintf("%d + days ago", day)), len(strconv.FormatBool(task.Done)))
	}
	return maxLength, nil
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
