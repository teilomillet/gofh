package main

import (
	"log"
	"strconv"

	"github.com/teilomillet/gofh"
)

var todos []string

const (
	todosEndpoint = "/todos"
	todoListID    = "#todo-list"
)

func main() {
	app := gofh.New()
	app.Get("/").Handle(handleHome)
	app.Post(todosEndpoint).Handle(handleTodoAction)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(app.Serve())
}

func handleHome(c *gofh.Context) gofh.Element {
	return renderFullPage()
}

func handleTodoAction(c *gofh.Context) gofh.Element {
	switch c.GetFormValue("action") {
	case "add":
		if title := c.GetFormValue("title"); title != "" {
			todos = append(todos, title)
		}
	case "toggle":
		if id, err := strconv.Atoi(c.GetFormValue("id")); err == nil && id < len(todos) {
			todos = append(todos[:id], todos[id+1:]...)
		}
	}
	return renderTodoList()
}

func renderFullPage() gofh.Element {
	return gofh.Div(
		gofh.H1("Todo List").Class("text-3xl font-bold mb-4 text-center"),
		gofh.Form(
			gofh.Input("text", "title").
				Attr("placeholder", "Enter a new todo").
				Class("input input-bordered w-full max-w-xs mr-2"),
			gofh.Button("Add").
				Attr("type", "submit").
				Class("btn btn-primary"),
		).HxPost(todosEndpoint+"?action=add").
			HxTarget(todoListID).
			HxSwap("outerHTML").
			Class("flex justify-center mb-4"),
		renderTodoList(),
	).Class("container mx-auto p-4 max-w-md")
}

func renderTodoList() gofh.Element {
	items := make([]gofh.Element, len(todos))
	for i, title := range todos {
		items[i] = gofh.Li(
			gofh.Input("checkbox", "").
				HxPost(todosEndpoint+"?action=toggle&id="+strconv.Itoa(i)).
				HxTarget(todoListID).
				HxSwap("outerHTML").
				HxTrigger("change").
				Class("checkbox checkbox-primary mr-2"),
			gofh.Span(title).Class("text-lg"),
		).Class("flex items-center mb-2")
	}
	return gofh.Ul(items...).Attr("id", "todo-list").Class("list-none")
}