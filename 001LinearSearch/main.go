package main

import (
	"fmt"
)

func search(data []int, target int) int {
	for i := 0; i < len(data); i++ {
		if data[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	data := []int{24, 18, 142, 9, 62, 56}
	res := search(data, 56)
	fmt.Println(res)
}
