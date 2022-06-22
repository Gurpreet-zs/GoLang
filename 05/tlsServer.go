package main

import (
	"crypto/tls"
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
	res.Header().Set("Content-Type", "application/json")

	switch r.URL.Path {
	case "/hello":
		fmt.Fprint(res, "Hello Interns!")
	case "/ping":
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(m)
	}

}

func main() {
	m := Message{Msg: "Pong"}
	mux := http.NewServeMux()
	mux.HandleFunc("/", m.ServeHttp)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	s := &http.Server{
		// Addr:         ":443",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	fmt.Println("Server listing on https")
	log.Fatal(s.ListenAndServeTLS("zop.local+3.pem", "zop.local+3-key.pem"))
}
