// internal/core/app.go

package core

import (
	"net/http"
	"strings"
)

type RouteHandler func(*Context) Element

type App struct {
	routes      map[string]map[string]RouteHandler
	templateDir string
}

type Context struct {
	Request *http.Request
	params  map[string]string
}

func NewApp() *App {
	return &App{
		routes: make(map[string]map[string]RouteHandler),
	}
}

// New RouteBuilder struct
type RouteBuilder struct {
	app    *App
	method string
	path   string
}

// New methods for fluent API
func (a *App) Get(path string) *RouteBuilder {
	return &RouteBuilder{app: a, method: "GET", path: path}
}

func (a *App) Post(path string) *RouteBuilder {
	return &RouteBuilder{app: a, method: "POST", path: path}
}

func (rb *RouteBuilder) Handle(handler RouteHandler) {
	rb.app.Route(rb.method, rb.path, handler)
}

// Existing Route method (kept for backwards compatibility)
func (a *App) Route(method, path string, handler RouteHandler) {
	if _, ok := a.routes[path]; !ok {
		a.routes[path] = make(map[string]RouteHandler)
	}
	a.routes[path][method] = handler
}

func (a *App) SetTemplateDir(dir string) {
	a.templateDir = dir
}

func (a *App) HandleRequest(w http.ResponseWriter, r *http.Request) Element {
	for routePath, methodHandlers := range a.routes {
		if handler, params := a.matchRoute(routePath, r.URL.Path, r.Method, methodHandlers); handler != nil {
			ctx := &Context{
				Request: r,
				params:  params,
			}
			return handler(ctx)
		}
	}
	return El("h1", "404 Not Found")
}

func (a *App) matchRoute(routePath, requestPath, method string, methodHandlers map[string]RouteHandler) (RouteHandler, map[string]string) {
	routeParts := strings.Split(routePath, "/")
	requestParts := strings.Split(requestPath, "/")

	if len(routeParts) != len(requestParts) {
		return nil, nil
	}

	params := make(map[string]string)
	for i, part := range routeParts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			paramName := strings.Trim(part, "{}")
			params[paramName] = requestParts[i]
		} else if part != requestParts[i] {
			return nil, nil
		}
	}

	if handler, ok := methodHandlers[method]; ok {
		return handler, params
	}

	return nil, nil
}

func (c *Context) GetFormValue(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) GetURLParam(key string) string {
	return c.params[key]
}

func (a *App) Serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		element := a.HandleRequest(w, r)
		RenderToResponse(w, element)
	})

	return http.ListenAndServe(":8080", nil)
}

