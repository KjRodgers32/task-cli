package data

import "time"

type Data struct {
	ID string
	Task string
	Created time.Time
	Done bool
}

type TestData struct {
	Data
	expected_days int
	expected_error error
}

type DataMethods interface {
	DaysBetweenTasks() int
}