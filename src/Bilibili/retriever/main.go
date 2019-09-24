package main

import (
	"Bilibili/retriever/Mock"
	"Bilibili/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.immooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{}
	fmt.Printf("%T %v\n", r, r)
	r = &real.Retreiever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
}
