package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{11, 12, 13, 14, 15}
	arr3 := []int{101, 102, 103, 104, 105}

	arr := [][]int{arr1, arr2, arr3}
	fmt.Println(arr[1][4])
}
