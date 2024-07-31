package main

import (
	"log"

	"github.com/teilomillet/gofasthtml"
)

func main() {
	app := gofasthtml.NewApp()

	app.Route("/", func() gofasthtml.Element {
		return gofasthtml.Div(
			gofasthtml.H1("Welcome to GoFastHTML"),
			gofasthtml.P("Hello World!").HX("get", "/change").HX("swap", "outerHTML"),
		)
	})

	app.Route("/change", func() gofasthtml.Element {
		return gofasthtml.P("Nice to be here!").HX("get", "/").HX("swap", "outerHTML")
	})

	log.Fatal(app.Serve(":8080"))
}
