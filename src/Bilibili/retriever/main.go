package main

import (
	"Bilibili/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.immooc.com")
}

func main() {
	var r Retriever
	r = real.Retreiever{}
	fmt.Println(download(r))
}
