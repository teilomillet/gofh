package main

import (
	"log"
	"os"

	"github.com/teilomillet/gofh"
	"github.com/teilomillet/gollm"
)

var llm gollm.LLM

func main() {
	var err error
	llm, err = gollm.NewLLM(
		gollm.SetProvider("openai"),
		gollm.SetModel("gpt-4o-mini"),
		gollm.SetAPIKey(os.Getenv("OPENAI_API_KEY")),
	)
	if err != nil {
		log.Fatal(err)
	}

	app := gofh.New()

	app.Get("/").Handle(func(c *gofh.Context) gofh.Element {
		return gofh.Div(
			gofh.Div().ID("chat"),
			gofh.Form(
				gofh.Input("text", "msg").Placeholder("Type a message"),
				gofh.Button("Send"),
			).Attr("hx-post", "/chat").Attr("hx-target", "#chat").Attr("hx-swap", "beforeend"),
		)
	})

	app.Post("/chat").Handle(func(c *gofh.Context) gofh.Element {
		msg := c.GetFormValue("msg")
		resp, err := llm.Generate(c.Request.Context(), gollm.NewPrompt(msg))
		if err != nil {
			resp = "Error: " + err.Error()
		}
		return gofh.Div(
			gofh.P("You: "+msg),
			gofh.P("AI: "+resp),
		)
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(app.Serve())
}

