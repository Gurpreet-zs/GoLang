package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ping struct {
	Message string `json:"message"`
}

func main() {

	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Interns!")
	})

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		var ping Ping
		ping.Message = "Pong"
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(ping)
	})

	fmt.Println("Server listening on 8000")
	http.ListenAndServe(":8000", nil)
}
