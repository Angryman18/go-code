package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// PROBLEM: Solved problem with another way;

func mockWait() time.Duration {
	wait := time.Second * time.Duration((rand.Intn(10)))
	time.Sleep(wait)
	return wait
}

func main() {

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func() {
		fmt.Println("Got the data from Passed Here")
		defer wg.Done()
		fmt.Println(mockWait())
		output := rand.Intn(1000)
		ch1 <- output
	}()
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
	}()

	wg.Wait()

}
