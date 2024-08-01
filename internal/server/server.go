// internal/server/server.go

package server

import (
	"fmt"
	"net/http"

	"github.com/teilomillet/gofh/internal/core"
)

type Server struct {
	addr      string
	staticDir string
}

func NewServer() *Server {
	return &Server{
		addr: ":8080", // Default port
	}
}

func (s *Server) SetAddr(addr string) {
	if addr != "" {
		s.addr = addr
	}
}

func (s *Server) SetStaticDir(dir string) {
	s.staticDir = dir
}

func (s *Server) Serve(app *core.App) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		element := app.HandleRequest(w, r)
		s.renderWithHTMX(w, element)
	})

	if s.staticDir != "" {
		fs := http.FileServer(http.Dir(s.staticDir))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	}

	return http.ListenAndServe(s.addr, nil)
}

func (s *Server) renderWithHTMX(w http.ResponseWriter, element core.Element) {
	w.Header().Set("Content-Type", "text/html")
	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>GoFH App</title>
    <script src="https://unpkg.com/htmx.org@1.9.2"></script>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://cdn.jsdelivr.net/npm/daisyui@3.1.6/dist/full.css" rel="stylesheet" type="text/css" />

</head>
<body>
    %s
</body>
</html>
`
	renderedElement := core.Render(element)
	fullHTML := fmt.Sprintf(htmlTemplate, renderedElement)
	w.Write([]byte(fullHTML))
}
