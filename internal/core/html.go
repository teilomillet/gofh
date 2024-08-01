package core

// Div creates a new div element
func Div(children ...Element) Element {
	return El("div", "", children...)
}

// P creates a new paragraph element
func P(content string) Element {
	return El("p", content)
}

// A creates a new anchor element
func A(href, content string) Element {
	return El("a", content).Attr("href", href)
}

// Span creates a new span element
func Span(content string) Element {
	return El("span", content)
}

// Input creates a new input element
func Input(type_, name string) Element {
	return El("input", "").Attr("type", type_).Attr("name", name)
}

// Form creates a new form element
func Form(children ...Element) Element {
	return El("form", "", children...)
}

// Button creates a new button element
func Button(content string) Element {
	return El("button", content)
}

// H1 creates a new h1 element
func H1(content string) Element {
	return El("h1", content)
}

// H2 creates a new h2 element
func H2(content string) Element {
	return El("h2", content)
}

// H3 creates a new h3 element
func H3(content string) Element {
	return El("h3", content)
}

// Ul creates a new unordered list element
func Ul(children ...Element) Element {
	return El("ul", "", children...)
}

// Li creates a new list item element
func Li(children ...Element) Element {
	return El("li", "", children...)
}

// Add more HTML element functions as needed...

