package main

import "fmt"

type MyError struct {
}

// implements error interface
func (MyError) Error() string {
	return "This is MyError"
}

func fn() (int, *MyError) {
	return 42, nil
}

func main() {
	var err error
	result, err := fn()
	if err != nil {
		fmt.Println("This should not be called")
	}
	fmt.Println(result, err)
}
