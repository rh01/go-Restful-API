package example

import (
	"net/http"
	"fmt"
	"html"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
