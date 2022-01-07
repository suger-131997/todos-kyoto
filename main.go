package main

import (
	"html/template"
	"net/http"

	"github.com/kyoto-framework/kyoto"
)

func ssatemplate(p kyoto.Page) *template.Template {
	return template.Must(template.New("SSA").Funcs(kyoto.TFuncMap()).ParseGlob("*.html"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", kyoto.PageHandler(&PageIndex{}))
	mux.HandleFunc("/todos", kyoto.PageHandler(&PageTodos{}))
	mux.HandleFunc("/game", kyoto.PageHandler(&PageGame{}))

	mux.HandleFunc("/ssa/", kyoto.SSAHandler(ssatemplate))

	http.ListenAndServe("localhost:9000", mux)
}