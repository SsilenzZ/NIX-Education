package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func responseWrite(i int) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	for true {
		bs := make([]byte, 512)
		n, err := resp.Body.Read(bs)
		ioutil.WriteFile("storage/posts/"+strconv.Itoa(i)+".txt", bs[:n], 0644)
		if n == 0 || err != nil {
			break
		}
	}
}

func main() {
	for i := 1; i <= 5; i++ {
		go responseWrite(i)
		time.Sleep(3 * time.Second)
	}
	for i := 1; i <= 5; i++ {
		content, err := ioutil.ReadFile("storage/posts/" + strconv.Itoa(i) + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(content))
	}
}
