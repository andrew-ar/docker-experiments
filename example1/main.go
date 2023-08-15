package main

import (
	"flag"
	"net/http"
)

func setupServer(listenAddr string) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(([]byte("pong")))
	})

	s := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}
	return &s
}

func main() {
	listenAddr := flag.String("listen-addr", "0.0.0.0:8080", "IP address and port number to listen on, in the format ADDRESS:PORT")
	flag.Parse()

	s := setupServer(*listenAddr)
	s.ListenAndServe()
}
