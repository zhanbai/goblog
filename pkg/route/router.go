package route

import (
	"github.com/gorilla/mux"
	"github.com/zhanbai/goblog/routes"
)

// Router 路由对象
var Router *mux.Router

// Initialize 初始化路由
func Initialize() {
	Router = mux.NewRouter()
	routes.RegisterWebRoutes(Router)
}
