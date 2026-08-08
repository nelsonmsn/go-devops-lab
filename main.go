package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go DevOps Lab@@@")
	})

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
