package main

import (
	"fmt"
	"sort"
)

type MyType struct {
	name string
	age  int
}

func main() {
	data := []MyType{
		MyType{
			name: "miku",
			age:  16,
		},
		MyType{
			name: "rin",
			age:  14,
		},
		MyType{
			name: "len",
			age:  14,
		},
		MyType{
			name: "luka",
			age:  20,
		},
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].name > data[j].name
	})
	fmt.Println(data)
}
