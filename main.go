package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhanbai/goblog/app/http/middlewares"
	"github.com/zhanbai/goblog/pkg/route"
)

var router *mux.Router

func main() {
	route.Initialize()
	router = route.Router

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
