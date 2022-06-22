package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{\"message\": \"Pong\"}")
	})

	fmt.Println("Server started and listening on 8000")
	http.ListenAndServe(":8000", nil)
}
