package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	// using the request "ctx"
	ctx := r.Context()
	log.Println("Request initiate")
	defer log.Println("Finish request")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Request handle successfully")
		w.Write([]byte("Request handle successfully"))
	case <-ctx.Done():
		log.Println("Canceled Request")
		http.Error(w, "Canceled Request", http.StatusRequestTimeout)
	}
}
