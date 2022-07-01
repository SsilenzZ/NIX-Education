package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	for true {
		bs := make([]byte, 512)
		n, err := resp.Body.Read(bs)
		fmt.Print(string(bs[:n]))
		if n == 0 || err != nil {
			break
		}
	}
}
