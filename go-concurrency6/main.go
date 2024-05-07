package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// PROBLEM: Spwan a x number of go routine and execute them in order and make sure each one of them are
// dependent to previous's output.

// like 3 api and 1st one's output is send to 2nd and 2nd's one are send to 3rd

func mockWait() time.Duration {
	wait := time.Second * time.Duration((rand.Intn(10)))
	time.Sleep(wait)
	return wait
}

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	wg := &sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())

	wg.Add(3)
	go func() {
		data := <-ch1
		fmt.Println("Got the data from ", data)
		defer wg.Done()
		defer close(ch1)
		fmt.Println(mockWait())
		output := rand.Intn(1000)
		ch2 <- output
	}()
	go func() {
		data := <-ch2
		fmt.Println("Got the data from ", data)
		defer wg.Done()
		defer close(ch2)
		fmt.Println(mockWait())
		output := rand.Intn(1000)
		ch3 <- output
	}()
	go func() {
		data := <-ch3
		fmt.Println("Got the data from ", data)
		defer wg.Done()
		defer close(ch3)
		fmt.Println(mockWait())
		// output := rand.Intn(1000)
		// ch2 <- output
	}()

	ch1 <- "My Payload"
	wg.Wait()

}
