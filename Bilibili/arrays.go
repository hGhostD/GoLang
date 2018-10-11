package main

import "fmt"

func printArray(arr [5]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8 ,10}
	fmt.Println(arr1, arr2, arr3)
	
	printArray(arr1)
	printArray(arr3)
	
}
