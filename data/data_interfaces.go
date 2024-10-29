package data

import (
	"errors"
	"math"
	"time"
)

func (d *Data) DaysBetweenTasks() (int, error) {
	timeInBetween := time.Since(d.Created)
	if timeInBetween.Hours() <= 0 {
		return 0, errors.New("you can't create a task from the future...unlesssssss...nevermind")
	} else {
		return int(math.Floor(timeInBetween.Hours()/24)) - 1, nil
	}
}

func (td *TestData) DaysBetweenTasks() (int, error) {
	timeInBetween := time.Since(td.Created)
	if timeInBetween.Hours() <= 0 {
		return 0, errors.New("you can't create a task from the future...unlesssssss...nevermind")
	} else {
		return int(math.Floor(timeInBetween.Hours()/24)) - 1, nil
	}
}

