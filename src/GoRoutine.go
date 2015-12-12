package main

import (
	. "fmt"
	"net/http"
	"io/ioutil"
	. "time"
)

func main() {

	url := []string{
		"http://jsonplaceholder.typicode.com/posts/1",
		"http://jsonplaceholder.typicode.com/posts/2",
		"http://jsonplaceholder.typicode.com/posts/3",
	}

	for _, u := range url {
		go InvokeUrl(u)
//		InvokeUrl(u)
	}

	Sleep(10 * Second)
}

func InvokeUrl(url string) {

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	Sleep(5 * Second) // Pretend doing something...

	Println(string(body))
}
