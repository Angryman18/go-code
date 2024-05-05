package main

import (
	"fmt"
	"sync"
)

// PROBLEM : Spwan Multiple Go Routines and send data to one channel and receive the data propery on the main
// thread.

func main() {

	ch := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 1; i <= 10; i++ { // spawning 10 go routines proceess
		wg.Add(1)
		go func(ch chan string, i int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- fmt.Sprintf("This is a data %v", i)
		}(ch, i, wg)
	}

	// main thread has blocking code so cannot wait for the
	// waitgroups to end there thats y spawn another routine to halt
	// the main thread and upon ending the wait close the channel since
	// the all the above routines are made Done() and channel is done sending
	// data
	go func() {
		wg.Wait()
		close(ch)
	}()

	// fmt.Println(<-ch) // receiving single data
	for el := range ch {
		fmt.Println(el)
	}

}
