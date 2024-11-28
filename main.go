package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target := "http://192.168.0.179" + r.URL.Path
		if r.URL.RawQuery != "" {
			target += "?" + r.URL.RawQuery
		}
		fmt.Println("r.URL.Path", r.URL.Path)
		http.Redirect(w, r, target, http.StatusTemporaryRedirect)
	})

	fmt.Println("Starting redirect server on :80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
