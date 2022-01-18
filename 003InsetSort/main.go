package main

import "fmt"

func insertSort(arr []int) []int {
	for i := range arr {
		p := i - 1
		data := arr[i]
		for p >= 0 && arr[p] > data {
			arr[p+1] = arr[p]
			p -= 1
		}
		arr[p+1] = data
	}
	return arr
}

func main() {
	array := []int{5, 7, 9, 4, 8, 3, 1, 2, 6, 3}
	res := insertSort(array)
	fmt.Println(res)
}
