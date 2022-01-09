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

	// ルーティングの設定
	mux.HandleFunc("/", kyoto.PageHandler(&PageTodos{}))

	// SSA用の設定
	mux.HandleFunc("/ssa/", kyoto.SSAHandler(ssatemplate))

	// サーバー起動
	http.ListenAndServe("localhost:9000", mux)
}
