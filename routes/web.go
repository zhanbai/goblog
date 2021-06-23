package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhanbai/goblog/app/http/controllers"
	"github.com/zhanbai/goblog/app/http/middlewares"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	// 静态页面
	pc := new(controllers.PagesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 中间件：强制内容类型为 HTML
	r.Use(middlewares.ForceHTML)
}
