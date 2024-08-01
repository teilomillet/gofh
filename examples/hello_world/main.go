// examples/hello_world/main.go

package main

import (
	"log"

	"github.com/teilomillet/gofh"
)

var isAlternateGreeting bool

func main() {
	// Create a new GoFH application with the server address
	app := gofh.New()

	// Define the route for the home page
	app.Get("/").Handle(func(c *gofh.Context) gofh.Element {
		return gofh.Div(
			gofh.H1("Welcome to GoFastHTML"),
			renderGreeting("Hello World!"),
		)
	})

	// Define the route for changing the greeting
	app.Get("/change").Handle(func(c *gofh.Context) gofh.Element {
		isAlternateGreeting = !isAlternateGreeting
		if isAlternateGreeting {
			return renderGreeting("Nice to be here!")
		}
		return renderGreeting("Hello World!")
	})

	// Start the server
	log.Fatal(app.Serve())
}

func renderGreeting(message string) gofh.Element {
	return gofh.P(message).
		HxGet("/change").
		HxSwap("outerHTML").
		Attr("id", "greeting")
}
