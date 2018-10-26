package container

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
	
	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Printf("没有 key")
	}

	delete(m, "name")
	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Printf("没有 key")
	}
}
