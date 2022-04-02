package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(what string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- what
		fmt.Println(i, what)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
