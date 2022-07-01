package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Posts struct {
	User  int    `json:"userId"`
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Comments struct {
	Post  int    `json:"postId"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Body  string `json:"body"`
}

func writeCommentsToDB(c []Comments, db *sql.DB, i int) {
	stmt, err := db.Prepare("INSERT INTO comments VALUES(?, ?, ?, ?, ?)")
	_, err = stmt.Exec(c[i].Post, c[i].Id, c[i].Name, c[i].Email, c[i].Body)
	if err != nil {
		return
	}
}

func writePostsToDB(i int, p []Posts, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO posts VALUES(?, ?, ?, ?)")
	_, err = stmt.Exec(p[i].User, p[i].Id, p[i].Title, p[i].Body)
	if err != nil {
		return
	}
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(p[i].Id))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var c []Comments
	if err := json.Unmarshal(body, &c); err != nil {
		log.Fatal(err)
	}
	for i := range c {
		go writeCommentsToDB(c, db, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:@/commentsdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=7")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var s []Posts
	if err := json.Unmarshal(body, &s); err != nil {
		log.Fatal(err)
	}
	for i := range s {
		go writePostsToDB(i, s, db)
		time.Sleep(1 * time.Second)
	}
}
