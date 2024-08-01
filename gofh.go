package gofh

import (
	"github.com/teilomillet/gofh/internal/core"
	"github.com/teilomillet/gofh/internal/server"
)

type Element = core.Element
type Context = core.Context

// App represents the main GoFH application
type App struct {
	core   *core.App
	server *server.Server
}

// New creates a new GoFH application with the given options
func New(options ...Option) *App {
	app := &App{
		core:   core.NewApp(),
		server: server.NewServer(),
	}

	for _, option := range options {
		option(app)
	}

	return app
}

// Text creates a text node Element
func Text(content string) Element {
	return core.El("", content)
}

// Get adds a new GET route to the application
func (a *App) Get(path string) *RouteBuilder {
	return &RouteBuilder{app: a, method: "GET", path: path}
}

// Post adds a new POST route to the application
func (a *App) Post(path string) *RouteBuilder {
	return &RouteBuilder{app: a, method: "POST", path: path}
}

// RouteBuilder helps build routes with a fluent API
type RouteBuilder struct {
	app    *App
	method string
	path   string
}

// Handle sets the handler for the route
func (rb *RouteBuilder) Handle(handler func(*Context) Element) {
	rb.app.core.Route(rb.method, rb.path, handler)
}

// Serve starts the GoFH application server
func (a *App) Serve() error {
	return a.server.Serve(a.core)
}

// Expose HTML element functions
var (
	Div    = core.Div
	P      = core.P
	A      = core.A
	Span   = core.Span
	Input  = core.Input
	Form   = core.Form
	Button = core.Button
	H1     = core.H1
	H2     = core.H2
	H3     = core.H3
	Ul     = core.Ul
	Li     = core.Li
	El     = core.El
	// Add other elements as needed
)
