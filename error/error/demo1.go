package error

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.When, e.What)
}

func Run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
