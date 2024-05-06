package main

import "fmt"

// PROBLEM: Spawn 2 go routine and send data to 2 channel and read these from simultaneously

func readFromChannel(ch1, ch2 <-chan string) {
	for {
		select {
		case val, ok := <-ch1:
			if !ok {
				return
			}
			fmt.Println("Ch1 -> ", val)
		case val, ok := <-ch2:
			if !ok {
				return
			}
			fmt.Println("Ch2 -> ", val)
		}
	}
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(ch chan<- string) {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%v", i)
		}
	}(ch1)

	go func(ch chan<- string) {
		defer close(ch)
		for i := 100; i > 90; i-- {
			ch <- fmt.Sprintf("%v", i)
		}
	}(ch2)

	readFromChannel(ch1, ch2)

}
