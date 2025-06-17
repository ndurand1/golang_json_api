package models

import "time"

type Todo struct {
	// idiomatic json: lowercase keys
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
