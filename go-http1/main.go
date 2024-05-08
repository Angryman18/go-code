package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseProcessor(payload *io.ReadCloser) (User, error) {
	data, err := io.ReadAll(*payload)
	if err != nil {
		return User{}, err
	}
	var user User
	json.Unmarshal(data, &user)
	return user, nil
}

func callAPI(id string, cha chan<- User) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%v", id)
	data, err := http.Get(url)

	if err != nil {
		log.Fatal("Something went wrong")
	}

	defer data.Body.Close()
	defer close(cha)
	resp, err := responseProcessor(&data.Body)
	if err != nil {
		fmt.Println("Error Occured")
		return
	}
	cha <- resp
}

type User struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Calling a http server")
	cha := make(chan User)

	go callAPI("3", cha)
	val := <-cha
	fmt.Println(val.Title)

}
