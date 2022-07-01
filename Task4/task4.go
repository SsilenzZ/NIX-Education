package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func responseGet(i int) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
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

func main() {
	for i := 1; i <= 100; i++ {
		go responseGet(i)
		time.Sleep(40 * time.Millisecond)
	}
}
