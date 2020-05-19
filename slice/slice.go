package main

import "fmt"

func main() {
	var ints = make([]int, 0) //空のsliceを作成
	// var ints = make([]int, 3)    // ints[0]-ints[2]まで0で初期化
	// var ints = make([]int, 0, 3) // len(ints) == 0 だが要素3個の容量を最初に確保する
	fmt.Println(ints)

	ints = append(ints, 42)
	fmt.Println(ints)
}
