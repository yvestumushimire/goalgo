package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	arr := []int{31, 41, 59, 26, 41, 58}
	result := InsertionSort(arr)
	elapsed := time.Since(start)
	fmt.Printf("results: %v took: %v \n", result, elapsed)
	result2 := InsertionSortReverse(arr)
	fmt.Printf("result2: %v took: %v \n", result2, elapsed)
}

func InsertionSort(arr []int) []int {
	for j, v := range arr {
		i := j - 1
		for i >= 0 && arr[i] > v {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = v
	}

	return arr
}

func InsertionSortReverse(arr []int) []int {
	for j, v := range arr {
		i := j - 1
		for i >= 0 && arr[i] < v {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = v
	}

	return arr
}

func mergeSort(arr []int) []int {

}
