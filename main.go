package main

import (
	"admin-panel/v1/app"
	"admin-panel/v1/pkg/middleware"
	"admin-panel/v1/pkg/route"
)

var (
	coreAPI *app.App
)

func main() {
	coreAPI = app.New()
	coreAPI.Middleware(
		middleware.GinMiddleware,
	)
	coreAPI.Route(
		route.ApiRoute,
	)
	coreAPI.Run(":8000")
}
