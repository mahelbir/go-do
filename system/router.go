package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-do/controllers"
	"go-do/middlewares"
	"go-do/utils"
)

// ================ ROUTES ===============

var routes = []route{
	{
		method:      http.MethodGet,
		path:        "/users/:id",
		middlewares: []gin.HandlerFunc{middlewares.AuthMiddleware()},
		handler:     controllers.GetUser,
	},
	//{
	//	path:        "/sub",
	//	middlewares: []gin.HandlerFunc{middlewares.AuthMiddleware()},
	//	routes: []route{
	//		{
	//			method:  http.MethodGet,
	//			path:    "/1",
	//			handler: controllers.TestController,
	//		},
	//	},
	//},
}

// ================ /ROUTES ===============

type route struct {
	method      string
	path        string
	middlewares []gin.HandlerFunc
	handler     gin.HandlerFunc
	routes      []route
}

func Router(router *gin.Engine) {
	setupRoutes(router, routes, "", nil)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.NotFoundResponse())
	})
}

func joinPaths(paths ...string) string {
	fullPath := ""
	for _, path := range paths {
		if path == "/" {
			continue
		}
		if fullPath == "" {
			fullPath = path
		} else {
			fullPath += "/" + path
		}
	}
	return fullPath
}

func setupRoutes(router *gin.Engine, routes []route, parentPath string, parentMiddlewares []gin.HandlerFunc) {
	for _, r := range routes {
		fullPath := joinPaths(parentPath, r.path)
		routeGroup := router.Group(fullPath)
		allMiddlewares := append(parentMiddlewares, r.middlewares...)
		routeGroup.Use(allMiddlewares...)
		if r.handler != nil && r.method != "" {
			switch r.method {
			case http.MethodGet:
				routeGroup.GET("", r.handler)
			case http.MethodPost:
				routeGroup.POST("", r.handler)
			case http.MethodPut:
				routeGroup.PUT("", r.handler)
			case http.MethodDelete:
				routeGroup.DELETE("", r.handler)
			case http.MethodPatch:
				routeGroup.PATCH("", r.handler)
			case http.MethodOptions:
				routeGroup.OPTIONS("", r.handler)
			case http.MethodHead:
				routeGroup.HEAD("", r.handler)
			default:
				routeGroup.Any("", r.handler)
			}
		}
		if len(r.routes) > 0 {
			setupRoutes(router, r.routes, fullPath, allMiddlewares)
		}
	}
}
