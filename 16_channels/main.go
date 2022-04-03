package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "hello"
	msg := <-c
	fmt.Println(msg)
}
