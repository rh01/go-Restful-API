package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	. "../models"
	)

func ToDoList(w http.ResponseWriter, r *http.Request)  {
	todos := Todos{
		Todo{Name: "write presentation"},
		Todo{Name: "Host meetup"},
	}
	// 序列化json格式
	json.NewEncoder(w).Encode(todos)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

