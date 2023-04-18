package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s at %v", e.what, e.when)
}

func run1() error {
	return &MyError{
		when: time.Now(),
		what: "Runtime error",
	}
}

func run2(i int) error {
	if i < 0 {
		return &MyError{time.Now(), "Entered value is negative"}
	} else {
		return nil
	}
}

func main() {
	if err := run1(); err != nil {
		fmt.Println(err)
	}
	if err := run2(-2); err != nil {
		fmt.Println(err)
	}

}
