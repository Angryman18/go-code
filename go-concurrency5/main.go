package main

import "fmt"

// PROBLEM: Spawn 2 go routine and send data to 2 channel and read these from simultaneously

func readFromChannel(ch1, ch2 <-chan string) {
	for { // infinite for loop
		select {
		case val, ok := <-ch1: // checking if data is coming from ch1
			if !ok {
				ch1 = nil // making the ch1 to nil if the channel is closed
			} else {
				fmt.Println("Ch1 -> ", val) // doing its operation
			}
		case val, ok := <-ch2: // checking if coming from ch2
			if !ok {
				ch2 = nil // making ch2 nil if the channel is closed
			} else {
				fmt.Println("Ch2 -> ", val) // doing its operation
			}
		}
		if ch1 == nil && ch2 == nil { // both channels are closed then break out of loop
			break
		}
	}
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	// 1. spawn this routine
	go func(ch chan<- string) { // 2. chan<- makes sure this is a sender channel
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("%v", i) // 3.sending data to ch1
		}
	}(ch1)

	// 1. spawn this process
	go func(ch chan<- string) { // 2. sender channle
		defer close(ch)
		for i := 100; i > 90; i-- {
			ch <- fmt.Sprintf("%v", i) // 3. sending to ch2
		}
	}(ch2)

	readFromChannel(ch1, ch2) // 4. function receiving data

}
