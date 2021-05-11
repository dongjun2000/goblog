package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	pc := new(controllers.PagesController)
	ac := new()

	// 静态页面
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	// 404
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
}
