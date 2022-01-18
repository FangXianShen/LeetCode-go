package main

import "fmt"

func selectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	var theArray = []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 17}
	fmt.Print("排序前")
	fmt.Println(theArray)
	fmt.Print("排序后")
	arrayResult := selectSort(theArray)
	fmt.Println(arrayResult)
}
