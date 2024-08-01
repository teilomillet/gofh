// render.go

package core

import (
	"net/http"
)

// Render converts an Element to its HTML string representation
func Render(e Element) string {
	return e.Render() // Remove html.EscapeString
}

// RenderToResponse writes the rendered Element to an http.ResponseWriter
func RenderToResponse(w http.ResponseWriter, e Element) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(Render(e)))
}
