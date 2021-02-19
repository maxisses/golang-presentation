package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // HL
		_, _ = w.Write([]byte("Hello CIC")) // HL
	})
	http.ListenAndServe(":80", nil) // HL
}
