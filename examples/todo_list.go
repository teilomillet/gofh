package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/teilomillet/gofasthtml"
)

type Todo struct {
	ID   int
	Text string
}

var todos []Todo
var nextID = 1

func main() {
	app := gofasthtml.NewApp()

	app.Route("/", func() gofasthtml.Element {
		return gofasthtml.Div(
			gofasthtml.H1("GoFastHTML Todo List"),
			todoForm(),
			todoList(),
		)
	})

	app.Route("/add-todo", func() gofasthtml.Element {
		text := gofasthtml.GetFormValue("todo-text")
		if text == "" {
			return gofasthtml.Error(errors.New("empty todo"), "Todo text cannot be empty", 400)
		}
		todos = append(todos, Todo{ID: nextID, Text: text})
		nextID++
		return todoList()
	})

	app.Route("/remove-todo", func() gofasthtml.Element {
		idStr := gofasthtml.GetFormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return gofasthtml.Error(err, "Invalid todo ID", 400)
		}
		found := false
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return gofasthtml.Error(errors.New("todo not found"), "Todo not found", 404)
		}
		return todoList()
	})

	log.Fatal(app.Serve(":8080"))
}

func todoForm() gofasthtml.Element {
	return gofasthtml.Form(
		gofasthtml.Input("text", "todo-text").With("placeholder", "Enter a new todo"),
		gofasthtml.Button("Add Todo"),
	).HX("post", "/add-todo").HX("target", "#todo-list").HX("swap", "outerHTML")
}

func todoList() gofasthtml.Element {
	items := make([]gofasthtml.Element, len(todos))
	for i, todo := range todos {
		items[i] = gofasthtml.Div(
			gofasthtml.Span(todo.Text),
			gofasthtml.Button("Remove").
				HX("post", "/remove-todo").
				HX("target", "#todo-list").
				HX("swap", "outerHTML").
				With("name", "id").
				With("value", strconv.Itoa(todo.ID)),
		)
	}
	return gofasthtml.Div(items...).With("id", "todo-list")
}

