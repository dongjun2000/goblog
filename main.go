package main

import (
	_ "github.com/go-sql-driver/mysql"
	"goblog/app/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"net/http"
)

func main() {
	database.Initialize()

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	// 中间件：强制内容类型为 HTML
	router.Use(middlewares.ForceHTML)

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
