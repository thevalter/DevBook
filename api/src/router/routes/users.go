package routes

import (
	"api/src/controllers"
	"net/http"
)

var UsersRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: controllers.GetUser,
		NeedAuth: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodDelete,
		Function: controllers.GetUserById,
		NeedAuth: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: false,
	},
}