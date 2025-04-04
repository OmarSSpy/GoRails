package main

import (
	"fmt"
	"gorails/core/router"
	"gorails/core/template"
	"net/http"
)

func main() {
	template.LoadTemplates()
	r := router.NewRouter()

	// Route with HTML rendering
	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		template.Render(w, "index.html", map[string]string{"Title": "Welcome to GoRails!"})
	})

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", r)
}
