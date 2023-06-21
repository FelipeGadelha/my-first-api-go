package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	case "PATCH":
	}
	fmt.Fprintf(w, "Hello Felipe! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
