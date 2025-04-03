package main

import (
	"fmt"
	"gorails/core/router"
	"net/http"
)

func main() {
	r := router.NewRouter()

	// Static route
	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		fmt.Fprintln(w, "Welcome to GoRails!")
	})

	// Dynamic route
	r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		fmt.Fprintf(w, "User ID: %s", params["id"])
	})

	http.ListenAndServe(":8000", r)
}
