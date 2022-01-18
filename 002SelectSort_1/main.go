package main

import "fmt"

func SelectSortMax(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		max := i
		for j := i + 1; j < length; j++ {
			if arr[max] < arr[j] {
				max = j
			}
		}
		arr[max], arr[i] = arr[i], arr[max]
	}
	return arr
}

func main() {
	var theArray = []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 17}
	arrayResult := SelectSortMax(theArray)
	fmt.Println(arrayResult)
}
