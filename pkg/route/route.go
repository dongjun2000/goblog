package route

import "github.com/gorilla/mux"

// Router 路由对象
var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
}

// 通过路由名称来获取URL
func Name2URL(routeName string, paris ...string) string {
	url, err := Router.Get(routeName).URL(paris...)
	if err != nil {
		//checkError(err)
		return ""
	}

	return url.String()
}
