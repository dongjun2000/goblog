package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

var router *mux.Router

// 设置路由实例，以供 Name2URL 等函数使用
func SetRoute(r *mux.Router) {
	router = r
}

// 通过路由名称来获取URL
func Name2URL(routeName string, paris ...string) string {
	url, err := router.Get(routeName).URL(paris...)
	if err != nil {
		//checkError(err)
		return ""
	}

	return url.String()
}

// 通过传参URL路由参数名称获取值
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}