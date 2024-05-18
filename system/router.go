package system

import (
	"github.com/gin-gonic/gin"
	"go-do/controllers"
	"go-do/middlewares"
	"net/http"
)

// ================ ROUTES ===============

var routes = []route{
	{
		path:    "/main",
		handler: controllers.TestController,
	},
	{
		path:        "/mid",
		middlewares: []gin.HandlerFunc{middlewares.AuthMiddleware()},
		handler:     controllers.TestController,
	},
	{
		path:        "/sub",
		middlewares: []gin.HandlerFunc{middlewares.AuthMiddleware()},
		routes: []route{
			{
				path:    "/1",
				handler: controllers.TestController,
			},
		},
	},
}

// ================ /ROUTES ===============

type route struct {
	path        string
	middlewares []gin.HandlerFunc
	handler     gin.HandlerFunc
	routes      []route
}

func Router(router *gin.Engine) {
	setupRoutes(router, routes, "", nil)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, NotFoundResponse())
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
		if r.handler != nil {
			routeGroup.Any("", r.handler)
		}
		if len(r.routes) > 0 {
			setupRoutes(router, r.routes, fullPath, allMiddlewares)
		}
	}
}
