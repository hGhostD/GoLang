package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variable() {
	a, b  := 3, 4
	var s = "abc"
	
	fmt.Println(a, b , s)
}

func euler() {
	fmt.Println(
		cmplx.Pow(math.E, 1i * math.Pi) + 1 )
	
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func main() {
	
	
	fmt.Printf("Hello world")
	variable()
	euler()
	triangle()
}
