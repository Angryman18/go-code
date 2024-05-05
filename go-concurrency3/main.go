package main

import "fmt"

// PROBLEM: Spawn a single go routine and send a tons of data and receive it to main routine without using
// wait groups.

func main() {

	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 1; i <= 100; i++ {
			ch <- fmt.Sprintf("Sending the data %v", i)
		}
	}()

	for el := range ch {
		fmt.Println(el)
	}

}
