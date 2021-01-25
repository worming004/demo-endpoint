package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/demo-endpoint/text", handleEndpoint)
	log.Fatal(http.ListenAndServe(":9090", mux))
}

func handleEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("/demo-endpoint/text")
	w.Write([]byte("ok from endpoint"))
}
