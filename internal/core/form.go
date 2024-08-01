// In gofh/form.go
package core

type FormBuilder struct {
	method  string
	action  string
	target  string
	swap    string
	inputs  []Element
	buttons []Element
}

func NewForm() *FormBuilder {
	return &FormBuilder{}
}

func (fb *FormBuilder) Post(action string) *FormBuilder {
	fb.method = "POST"
	fb.action = action
	return fb
}

func (fb *FormBuilder) Target(target string) *FormBuilder {
	fb.target = target
	return fb
}

func (fb *FormBuilder) Swap(swap string) *FormBuilder {
	fb.swap = swap
	return fb
}

func (fb *FormBuilder) WithInput(name, placeholder string) *FormBuilder {
	fb.inputs = append(fb.inputs, Input(name, placeholder))
	return fb
}

func (fb *FormBuilder) WithButton(text string) *FormBuilder {
	fb.buttons = append(fb.buttons, Button(text))
	return fb
}

func (fb *FormBuilder) Build() Element {
	form := Form(append(fb.inputs, fb.buttons...)...)
	form.Attr("method", fb.method)
	form.Attr("action", fb.action)
	if fb.target != "" {
		form.Attr("hx-target", fb.target)
	}
	if fb.swap != "" {
		form.Attr("hx-swap", fb.swap)
	}
	return form
}
