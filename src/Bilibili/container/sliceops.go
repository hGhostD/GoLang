package container

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n",len(s), cap(s));
}

func main() {
	var s []int // Zero value for sli e is nil

	for i := 0; i < 100 ; i++  {
		s = append(s, 2 * i + 1)
		printSlice(s)
	}

}
