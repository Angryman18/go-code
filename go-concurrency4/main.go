package main

import (
	"fmt"
)

// PROBLEM: Spwan x number go routine sending data to x number channel and handle it on main thread without
// using wait groups

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		defer close(ch1)
		ch1 <- "Hello World 1"
	}()
	go func() {
		defer close(ch2)
		ch2 <- "Hello World 2"
	}()
	go func() {
		defer close(ch3)
		ch3 <- "Hello World 3"
	}()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println(<-ch3)

}
