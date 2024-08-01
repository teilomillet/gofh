// internal/core/element.go

package core

import (
	"fmt"
)

type Element interface {
	Render() string
	Attr(key, value string) Element
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

func (e *BaseElement) SetAttribute(key, value string) {
	if e.Attributes == nil {
		e.Attributes = make(map[string]string)
	}
	e.Attributes[key] = value
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
	e.SetAttribute(key, value)
	return e
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

