package basic

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

// 13 / 3 = 4 ... 1
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
	
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args " + "%d, %d\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(a ...int) int {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	return sum
}
func main() {
	fmt.Println(eval(3, 4, "*"))
	q, r := div(13, 4)
	fmt.Println(q, r)
	//fmt.Println(apply(pow, 3 ,4))
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b) ))
		}, 3 ,4))
	fmt.Println(sum(1, 2, 3, 4, 5))
}
