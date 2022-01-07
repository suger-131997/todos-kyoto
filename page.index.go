package main

import (
	"html/template"

	"github.com/kyoto-framework/kyoto"
)

type PageIndex struct {
	Counter kyoto.Component
}

func (*PageIndex) Template() *template.Template {
	return template.Must(template.New("page.index.html").Funcs(kyoto.TFuncMap()).ParseGlob("*.html"))
}

func (p *PageIndex) Init() {
	p.Counter = kyoto.RegC(p, &ComponentCounter{})
}
