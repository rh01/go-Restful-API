package models

import "time"

type Todo struct {
	Name     string    `json:"name"`
	Complete bool      `json:"complete"`
	Due      time.Time `json:"due"`
}

type Todos []Todo
