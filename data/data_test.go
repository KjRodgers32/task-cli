package data

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestDaysBetweenTasks(t *testing.T) {
	testCases := []TestData{
		{
			Data: Data{
				ID:      "1",
				Task:    "Move furniture",
				Created: time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC),
				Done:    false,
			},
			expected_days: 16,
		},
		{
			Data: Data{
				ID:      "2",
				Task:    "Walk Dog",
				Created: time.Date(2024, 9, 12, 0, 0, 0, 0, time.UTC),
				Done:    false,
			},
			expected_days: 46,
		},
		{
			Data: Data{
				ID:      "3",
				Task:    "Run a mile",
				Created: time.Date(2023, 10, 28, 0, 0, 0, 0, time.UTC),
				Done:    false,
			},
			expected_days: 366,
		},
		{
			Data: Data{
				ID:      "4",
				Task:    "Run two mile",
				Created: time.Date(2024, 10, 28, 0, 0, 0, 0, time.UTC),
				Done:    false,
			},
			expected_days: 0,
		},
		{
			Data: Data{
				ID:      "5",
				Task:    "Buy groceries",
				Created: time.Date(2025, 10, 12, 0, 0, 0, 0, time.UTC),
				Done:    false,
			},
			expected_error: errors.New("you can't create a task from the future...unlesssssss...nevermind"),
		},
	}

	pass := 0
	fail := 0

	for _, test := range testCases {
		actualDays, err := test.DaysBetweenTasks()
		fmt.Println("---------------------------------------------")
		if test.expected_error != nil {
			if err == nil || test.expected_error.Error() != err.Error() {
				t.Errorf(`
ID:             %s
Task:           %s
Created:        %v
Done:           %v
Expected Days:  %d
Actual Days:    %d
Expected Error: %v
Actual Erorr:   %v`,
					test.ID, test.Task, test.Created, test.Done, test.expected_days, actualDays, test.expected_error, err)
				fmt.Println("\nTest Failed")
				fail++
			} else {
				fmt.Printf(`
ID:             %s
Task:           %s
Created:        %v
Done:           %v
Expected Days:  %d
Actual Days:    %d
Expected Error: %v
Actual Erorr:   %v`,
					test.ID, test.Task, test.Created, test.Done, test.expected_days, actualDays, test.expected_error, err)
				fmt.Println("\nTest Passed")
				pass++
			}
		} else if test.expected_error == nil && test.expected_days == actualDays {
			fmt.Printf(`
ID:             %s
Task:           %s
Created:        %v
Done:           %v
Expected Days:  %d
Actual Days:    %d
Expected Error: %v
Actual Erorr:   %v`,
				test.ID, test.Task, test.Created, test.Done, test.expected_days, actualDays, test.expected_error, err)
			fmt.Println("\nTest Passed")
			pass++
		} else {
			t.Errorf(`
ID:             %s
Task:           %s
Created:        %v
Done:           %v
Expected Days:  %d
Actual Days:    %d
Expected Error: %v
Actual Erorr:   %v`,
				test.ID, test.Task, test.Created, test.Done, test.expected_days, actualDays, test.expected_error, err)
			fmt.Println("\nTest Failed")
			fail++
		}
	}

	fmt.Println("\n---------------------------------------------")
	if fail > 0 {
		fmt.Println("Test Failed")
	} else {
		fmt.Println("Test Passed")
	}

	fmt.Printf("Passed: %d  Failed: %d\n", pass, fail)

}

