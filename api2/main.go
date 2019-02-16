package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"message\": \"this is a message from API2\"}")
	})

	http.ListenAndServe(":8002", nil)
}
