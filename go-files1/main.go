package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./text.txt")
	if err != nil {
		fmt.Println("Invalid File Reading")
	}
	fileInfo, _ := file.Stat()
	b := make([]byte, fileInfo.Size())
	file.Read(b)
	fmt.Println(string(b))
}
