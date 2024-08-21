package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Todo struct {
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

		i := len(todos)
		if i == 0 {
			todos = make(map[int]Todo)
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
	var message string

	if id, err := strconv.Atoi(r.PathValue("id")); err != nil {
		message = "Malformed request; try again."
	} else {
		todo := todos[id]
		todo.Done = !todo.Done
		todos[id] = todo

		if todos[id].Done == true {
			message = fmt.Sprintf("Todo '%s' marked as done.", todo.Name)
		} else {
			message = fmt.Sprintf("Todo '%s' marked as todo.", todo.Name)
		}
	}

	pageData := map[string]any{
		"Title":   "Todo list",
		"Todos":   todos,
		"Message": message,
	}
	app.render(w, "todos", pageData, http.StatusOK)
}
