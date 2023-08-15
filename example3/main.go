package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"net/http"
	"os"
)

func setupServer(listenAddr string) *http.Server {
	mux := http.NewServeMux()

	secret1 := os.Getenv("SECRET1")
	secret2 := os.Getenv("SECRET2")

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		// Hash the query parameters together with the secrets and return the base64-encoded
		// SHA-256 hash in the response body.
		buffer := new(bytes.Buffer)
		buffer.Write([]byte(r.URL.RawQuery))
		buffer.Write([]byte(secret1))
		buffer.Write([]byte(secret2))

		hash := sha256.Sum256(buffer.Bytes())

		hashString := base64.StdEncoding.EncodeToString(hash[:])
		w.Write([]byte(hashString))
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
