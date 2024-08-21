package main

import (
	"net/http"
	"strconv"
)

type Todo struct {
	ID   int
	Name string
	Done bool
}

// Faux database.
var todos map[int]Todo

func (app *Application) Todos(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Title": "Todo list",
		"Todos": todos,
	}
	app.render(w, "todos", pageData, http.StatusOK)
}

func (app *Application) TodosAdd(w http.ResponseWriter, r *http.Request) {
	var message string

	r.ParseForm()

	if !r.Form.Has("name") || r.FormValue("name") == "" {
		message = "Malformed request; try again."
	} else {
		newTodo := Todo{
			Name: r.FormValue("name"),
			Done: false,
		}

		id := len(todos)
		if id == 0 {
			newTodo.ID = id
			todos = make(map[int]Todo)
		} else {
			newTodo.ID = id
		}
		todos[len(todos)] = newTodo

		message = "New todo added."
	}

	pageData := map[string]any{
		"Title":   "Todo list",
		"Todos":   todos,
		"Message": message,
	}
	app.render(w, "todos", pageData, http.StatusOK)
}

func (app *Application) TodosToggle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		return
	}

	todo := todos[id]
	todo.Done = !todo.Done
	todos[id] = todo

	pageData := map[string]any{
		"ID":   id,
		"Name": todo.Name,
		"Done": todo.Done,
	}
	app.render(w, "todo", pageData, http.StatusOK)
}
