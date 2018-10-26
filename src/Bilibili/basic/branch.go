package basic

import (
	"io/ioutil"
	"fmt"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	const filename = "abc.txt"
	if constents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", constents)
	}
	
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(92),
		grade(99),
		grade(100),
		//grade(-3),
		//grade(120),
	)
	
	a, b := 3, 43
	swap(&a, &b)
	fmt.Println(a, b)
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
