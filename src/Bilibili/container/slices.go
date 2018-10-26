package main

import "fmt"

func updateSlieces(arr []int) {
	arr[0] = 190

}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	fmt.Println(arr[:2])
	fmt.Println(arr[2:])
	fmt.Println(arr[2:6])
	fmt.Println(arr)

	s1 := arr[2:]
	s2 := arr[:]

	updateSlieces(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	updateSlieces(s2)
	fmt.Println(s2)
	fmt.Println(arr)
}
