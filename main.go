package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// show json response
		w.Write([]byte(`{"message": "Hello World"}`))
	})

	http.ListenAndServe(":8080", nil)
}
