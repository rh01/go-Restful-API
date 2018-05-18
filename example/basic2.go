package example

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/router/{router_id}", RouterId)
	log.Fatal(http.ListenAndServe(":8080", router))
}
func RouterId(w	 http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	router_id := vars["router_id"]

	fmt.Fprintln(w, "/router/" + router_id)
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Index dfdf")
}
