# gofh

Gofh -- GoFastHTML is a lightweight, fast web framework for Go, designed to create interactive web applications with minimal code, inspired by [FastHTML](https://github.com/AnswerDotAI/fasthtml). Gofh combines the speed of Go with the interactivity of HTMX to deliver powerful web experiences.


## Features

- Simple and expressive API
- Built-in HTMX integration for dynamic content
- Lightweight and fast
- Easy to learn for Go developers

## Getting Started

These instructions will get you a copy of gofh up and running on your local machine for development and testing purposes.

### Prerequisites

To use gofh, you need to have Go installed on your system. gofh is compatible with Go 1.16 and later.

### Installation

Install gofh using `go get`:

```bash
go get github.com/teilomillet/gofh
```

## Quick Start

Here's a simple example to get you started with gofh:

```go
package main

import (
	"log"

	"github.com/teilomillet/gofh"
)

func main() {
	app := gofh.New()

	app.Get("/").Handle(func(c *gofh.Context) gofh.Element {
		return gofh.Div(
			gofh.H1("Welcome to gofh"),
			gofh.P("Hello World!").HxGet("/greet").HxSwap("outerHTML"),
		)
	})

	app.Get("/greet").Handle(func(c *gofh.Context) gofh.Element {
		return gofh.P("Nice to meet you!")
	})

	log.Fatal(app.Serve())
}
```

Run this example and visit http://localhost:8080 in your browser to see it in action.

## Usage

gofh uses a simple, fluent API for defining routes and creating HTML elements:

```go

app := gofh.New()

app.Get("/").Handle(func(c *gofh.Context) gofh.Element {
	return gofh.Div(
		gofh.H1("Todo List"),
		gofh.Ul(
			gofh.Li("Item 1"),
			gofh.Li("Item 2"),
		),
	)
})
```

For more detailed usage instructions and examples, please refer to the documentation (Note: This link is a placeholder and may not be active yet).

## Contributing

We welcome contributions to gofh! Please feel free to submit issues, fork the repository and send pull requests!

## License

This project is licensed under the MIT License - see the LICENSE file for details.
Acknowledgments

- Inspired by FastHTML and other modern web frameworks
- Thanks to the Go community for their excellent tools and libraries
