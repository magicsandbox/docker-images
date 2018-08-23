package main

import (
	"fmt"
	"log"
	"net/http"
)

func makeResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Custom response from a go server\n")
}

func main() {
	http.HandleFunc("/", makeResponse)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
