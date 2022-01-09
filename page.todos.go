package main

import (
	"html/template"

	"github.com/kyoto-framework/kyoto"
)

type PageTodos struct {
	Todos kyoto.Component
}

func (*PageTodos) Template() *template.Template {
	return template.Must(template.New("page.todos.html").Funcs(kyoto.TFuncMap()).ParseGlob("*.html"))
}

func (p *PageTodos) Init() {
	//コンポーネントの初期化
	p.Todos = kyoto.RegC(p, &ComponentTodos{})
}
