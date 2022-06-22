package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Msg string `json:"message"`
}

func (m Message) ServeHttp(res http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		fmt.Println(res, "Hello Interns!")
	case "/ping":
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(m)
	}

}

func main() {
	m := Message{Msg: "Pong"}
	mux := http.NewServeMux()
	mux.HandleFunc("/", m.ServeHttp)

	s := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server listing on 8000")
	log.Fatal(s.ListenAndServe())
}
