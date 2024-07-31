package gofasthtml

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

// AppError represents an application-specific error
type AppError struct {
	Error   error
	Message string
	Code    int
	Stack   string
}

// NewAppError creates a new AppError
func NewAppError(err error, message string, code int) *AppError {
	return &AppError{
		Error:   err,
		Message: message,
		Code:    code,
		Stack:   string(debug.Stack()),
	}
}

// Element represents an HTML element or an error
type Element struct {
	tag        string
	content    string
	children   []Element
	attributes map[string]string
	err        *AppError
}

// Render converts an Element to its HTML string representation
func (e Element) Render() string {
	if e.err != nil {
		// Render error message
		return fmt.Sprintf("<div class='error'>Error: %s</div>", e.err.Message)
	}

	attrs := ""
	for k, v := range e.attributes {
		attrs += fmt.Sprintf(` %s="%s"`, k, v)
	}

	childContent := ""
	for _, child := range e.children {
		childContent += child.Render()
	}

	if e.content == "" && len(e.children) == 0 {
		return fmt.Sprintf("<%s%s />", e.tag, attrs)
	}
	return fmt.Sprintf("<%s%s>%s%s</%s>", e.tag, attrs, e.content, childContent, e.tag)
}

// With adds an attribute to an Element
func (e Element) With(key, value string) Element {
	if e.attributes == nil {
		e.attributes = make(map[string]string)
	}
	e.attributes[key] = value
	return e
}

// Add adds child elements to an Element
func (e Element) Add(children ...Element) Element {
	e.children = append(e.children, children...)
	return e
}

// HX adds an HTMX attribute to an Element
func (e Element) HX(key, value string) Element {
	return e.With("hx-"+key, value)
}

// El creates a new Element
func El(tag string, content string, children ...Element) Element {
	return Element{tag: tag, content: content, children: children}
}

// Error creates an error Element
func Error(err error, message string, code int) Element {
	return Element{err: NewAppError(err, message, code)}
}

// Common HTML elements
func Div(children ...Element) Element  { return El("div", "", children...) }
func P(content string) Element         { return El("p", content) }
func A(href, content string) Element   { return El("a", content).With("href", href) }
func Span(content string) Element      { return El("span", content) }
func Input(type_, name string) Element { return El("input", "").With("type", type_).With("name", name) }
func Form(children ...Element) Element { return El("form", "", children...) }
func Button(content string) Element    { return El("button", content) }
func H1(content string) Element        { return El("h1", content) }
func H2(content string) Element        { return El("h2", content) }
func H3(content string) Element        { return El("h3", content) }

// App represents the GoFastHTML application
type App struct {
	routes map[string]func() Element
}

// NewApp creates a new GoFastHTML application
func NewApp() *App {
	return &App{routes: make(map[string]func() Element)}
}

// Route adds a new route to the application
func (a *App) Route(path string, handler func() Element) {
	a.routes[path] = handler
}

var currentRequest *http.Request

// GetFormValue retrieves a form value from the current request
func GetFormValue(key string) string {
	if currentRequest != nil {
		return currentRequest.FormValue(key)
	}
	return ""
}

// ServeHTTP implements the http.Handler interface
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currentRequest = r
	r.ParseForm()

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n%s", r, debug.Stack())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	handler, ok := a.routes[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	element := handler()
	if element.err != nil {
		log.Printf("Error in route %s: %v\n%s", r.URL.Path, element.err.Error, element.err.Stack)
		w.WriteHeader(element.err.Code)
	}

	if strings.HasPrefix(r.Header.Get("HX-Request"), "true") {
		fmt.Fprint(w, element.Render())
	} else {
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>GoFastHTML page</title>
				<script src="https://unpkg.com/htmx.org@1.9.6"></script>
				<style>
					.error { color: red; border: 1px solid red; padding: 10px; margin: 10px 0; }
				</style>
			</head>
			<body>
				%s
			</body>
			</html>
		`, element.Render())
	}
}

// Serve starts the GoFastHTML application server
func (a *App) Serve(addr string) error {
	return http.ListenAndServe(addr, a)
}

