package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/demo-endpoint/text", handleEndpoint)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	log.Fatal(http.ListenAndServe(":9090", mux))
}

func handleEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("/demo-endpoint/text")
	w.Header().Add("custom-header", "custom-value")
	w.Write([]byte("ok from endpoint"))
}
