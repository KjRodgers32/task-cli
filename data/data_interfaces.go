package data

import (
	"time"
	"math"
	"errors"
)

func (d *Data) DaysBetweenTasks() (int, error) {
	timeInBetween := time.Since(d.Created)
	if timeInBetween.Hours() <= 24.0 {
		return 0, errors.New("you can't create a task from the future... unlesssssss... nevermind")
	} else {
		return int(math.Floor(timeInBetween.Hours() / 24)), nil
	}
}

func (td *TestData) DaysBetweenTasks() (int, error) {
	timeInBetween := time.Since(td.Created)
	if timeInBetween.Hours() <= 24.0 {
		return 0, errors.New("you can't create a task from the future... unlesssssss... nevermind")
	} else {
		return int(math.Floor(timeInBetween.Hours() / 24)), nil
	}
}