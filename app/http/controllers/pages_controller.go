package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

// PagesController 处理静态页面
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

// About 关于我们页面
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/views/about.gohtml")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

// NotFound 404 页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>如有疑惑，请联系我们。")
}
