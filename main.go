package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Can't read the body.", 400)
			return
		}
		fmt.Fprintf(w, "Hello, %s", d)
	})
	http.ListenAndServe(":9090", nil)
}
