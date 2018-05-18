package routers

import (
	"net/http"
	"github.com/gorilla/mux"
	. "../controllers"
)

// 构造route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router

}

var routes = Routes{
	Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: Index},
	Route{Name: "ToDoList", Method: "GET", Pattern: "/todolist", HandlerFunc: ToDoList},
}
