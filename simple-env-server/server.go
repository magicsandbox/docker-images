package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func makeResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, os.Getenv("HOSTNAME"))

}

func main() {
	http.HandleFunc("/", makeResponse)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
