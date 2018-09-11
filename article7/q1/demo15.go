package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	fmt.Printf("The length of s1:%d\n", len(s1))
	fmt.Printf("The capacity of s1:%d\n", cap(s1))
	
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2:%d\n", len(s2))
	fmt.Printf("The capacity of s2:%d\n", cap(s2))
	
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("%d %d %d", len(s4), cap(s4), s4)
	
}
