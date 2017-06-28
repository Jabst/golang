package main

import (
	"net/http"

	"./handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetAllUsers",
		"GET",
		"/all",
		handlers.GetAllUsers,
	},
	Route{
		"GetUserByID",
		"GET",
		"/userbyid",
		handlers.GetUserByID,
	},
	Route{
		"GetUserByName",
		"GET",
		"/userbyname",
		handlers.GetUserByName,
	},
	Route{
		"PostUser",
		"POST",
		"/user",
		handlers.PostUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/user",
		handlers.DeleteUser,
	},
	Route{
		"GetTaskByUser",
		"GET",
		"/taskuser",
		handlers.GetTaskByUser,
	},
	Route{
		"LoginUser",
		"POST",
		"/login",
		handlers.LoginUser,
	},
	Route{
		"CreateTask",
		"POST",
		"/task",
		handlers.PostTask,
	},
	Route{
		"DeleteTask",
		"DELETE",
		"/task",
		handlers.DeleteTask,
	},
}
