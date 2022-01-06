package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my bad site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Printf(err.Error())
	}
}
