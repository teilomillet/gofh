// internal/core/element.go

package core

import (
	"fmt"
	"strings"
)

type Element interface {
	Render() string
	Attr(key, value string) Element
	ID(id string) Element
	Class(classes ...string) Element
	Href(url string) Element
	Src(url string) Element
	Type(t string) Element
	Value(v string) Element
	Placeholder(p string) Element
	HxGet(url string) Element
	HxPost(url string) Element
	HxTrigger(event string) Element
	HxTarget(target string) Element
	HxSwap(strategy string) Element
}

type BaseElement struct {
	Tag        string
	Attributes map[string]string
	Children   []Element
	Content    string
}

func (e *BaseElement) Render() string {
	attrs := ""
	for k, v := range e.Attributes {
		attrs += fmt.Sprintf(` %s="%s"`, k, v)
	}

	childContent := ""
	for _, child := range e.Children {
		childContent += child.Render()
	}

	if e.Content == "" && len(e.Children) == 0 {
		return fmt.Sprintf("<%s%s />", e.Tag, attrs)
	}
	return fmt.Sprintf("<%s%s>%s%s</%s>", e.Tag, attrs, e.Content, childContent, e.Tag)
}

func (e *BaseElement) Attr(key, value string) Element {
	if e.Attributes == nil {
		e.Attributes = make(map[string]string)
	}
	e.Attributes[key] = value
	return e
}

func (e *BaseElement) ID(id string) Element {
	return e.Attr("id", id)
}

func (e *BaseElement) Class(classes ...string) Element {
	currentClass := e.Attributes["class"]
	if currentClass != "" {
		currentClass += " "
	}
	currentClass += strings.Join(classes, " ")
	return e.Attr("class", currentClass)
}

func (e *BaseElement) Href(url string) Element {
	return e.Attr("href", url)
}

func (e *BaseElement) Src(url string) Element {
	return e.Attr("src", url)
}

func (e *BaseElement) Type(t string) Element {
	return e.Attr("type", t)
}

func (e *BaseElement) Value(v string) Element {
	return e.Attr("value", v)
}

func (e *BaseElement) Placeholder(p string) Element {
	return e.Attr("placeholder", p)
}

func (e *BaseElement) HxGet(url string) Element {
	return e.Attr("hx-get", url)
}

func (e *BaseElement) HxPost(url string) Element {
	return e.Attr("hx-post", url)
}

func (e *BaseElement) HxTrigger(event string) Element {
	return e.Attr("hx-trigger", event)
}

func (e *BaseElement) HxTarget(target string) Element {
	return e.Attr("hx-target", target)
}

func (e *BaseElement) HxSwap(strategy string) Element {
	return e.Attr("hx-swap", strategy)
}

// El creates a new Element
func El(tag string, content string, children ...Element) Element {
	return &BaseElement{Tag: tag, Content: content, Children: children}
}
