package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	
	m2 := make(map[string]int)
	
	var m3 map[string]int
	fmt.Println(m, m2, m3)
	
	for k,v := range m {
		fmt.Println(k, v)
	}
	
	courseName, ok := m["course"]
	if couseName, ok := m["course"]; ok {
		fmt..
	}
	fmt.Println(len(courseName), ok)
}
