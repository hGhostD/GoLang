package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes 我爱ni!" // UTF-8
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	
	for i, v := range s {
		fmt.Printf("(%d %X)", i, v)
	}
	fmt.Println()
	
	fmt.Println(utf8.RuneCountInString(s))
	
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
	}
	fmt.Println()
}
