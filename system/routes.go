package system

import (
	"github.com/gin-gonic/gin"
	"go-do/controllers"
	"go-do/middlewares"
	"net/http"
)

var routes = []Route{
	{
		method:  http.MethodPost,
		path:    "/auth/login",
		handler: controllers.Login,
	},
	{
		method: http.MethodGet,
		path:   "/users/:id",
		middlewares: []gin.HandlerFunc{
			middlewares.Auth(),
			middlewares.Admin(),
			middlewares.ParamID(),
		},
		handler: controllers.GetUser,
	},
	{
		path:        "/todo_list",
		middlewares: []gin.HandlerFunc{middlewares.Auth()},
		routes: []Route{
			{
				method:  http.MethodPost,
				handler: controllers.CreateTodoList,
			},
			{
				method:  http.MethodGet,
				handler: controllers.ListTodoList,
			},
			{
				path:        "/:id",
				middlewares: []gin.HandlerFunc{middlewares.ParamID(), middlewares.TodoListAccess()},
				routes: []Route{
					{
						method:  http.MethodGet,
						handler: controllers.GetTodoList,
					},
					{
						method:  http.MethodPatch,
						handler: controllers.UpdateTodoList,
					},
					{
						method:  http.MethodDelete,
						handler: controllers.DeleteTodoList,
					},
				},
			},
		},
	},
}
