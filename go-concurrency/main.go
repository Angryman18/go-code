package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

// PROBLEM: We have a slice of x number of urls from where we have to download files. So download data from the
// url 4 at a time.

func createMockUrls(count int) []string {
	var output = []string{}

	for i := 1; i <= count; i++ {
		output = append(output, fmt.Sprint(uuid.NewString()))
	}
	return output
}

func downloadFile(ch chan file, wg *sync.WaitGroup) {
	defer wg.Done()
	for el := range ch { // receiving data from the channel
		fmt.Println(el.fileno, " -> We are Downloading from ", el.fileurl)
		time.Sleep(time.Second * time.Duration(rand.Intn(5))) // mimicking the download by waitng for a random sec
	}
}

type file struct {
	fileno  int
	fileurl string
}

func main() {

	data := createMockUrls(20) // assuming 20 urls
	concurrency := 4           // will be downloading 4 concurrently;
	ch := make(chan file, 4)   // buffer is 4
	wg := sync.WaitGroup{}

	// spawing 4 process at once
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go downloadFile(ch, &wg)
	}

	go func() {
		defer close(ch) // upon done sending we can close the channel
		for idx, el := range data {
			fmt.Println("<--- Pushing into the channel ", el)
			ch <- file{fileno: idx, fileurl: el} // sending the data to the channel
		}
	}()
	wg.Wait()

}

// NOTE:
// SENDING THE data from the main routine will not be reliable bcz once all the data are sent to the channel
// then main routine has no blocker code will exit out and thus downloading process will halt and since the buffer
// limit is 5 so there always be 5 url in the channel which will never be consumed bcz of the main routine has
// no blocker and program will exit out from the main routine
