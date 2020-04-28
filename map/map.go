package main

import "fmt"

func main() {
	values := make(map[string]int, 0)
	values["a"] = 42
	values["b"] = 43

	fmt.Println(values)
	fmt.Println(values["a"])
	fmt.Println(values["b"])
	fmt.Println(values["c"])
}
