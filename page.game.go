package main

import (
	"html/template"

	"github.com/kyoto-framework/kyoto"
)

type PageGame struct {
	Counter kyoto.Component
}

func (*PageGame) Template() *template.Template {
	return template.Must(template.New("page.game.html").Funcs(kyoto.TFuncMap()).ParseGlob("*.html"))
}

func (p *PageGame) Init() {
	p.Counter = kyoto.RegC(p, &ComponentCounter{})
}
