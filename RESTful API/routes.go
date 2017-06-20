package main

import (
	"net/http"
	"./Users"
)

type Route struct {
	Name		 string
	Method		 string
	Pattern		 string
	HandlerFunc	 http.HandlerFunc
}


type Routes []Route

var routes = Routes{
	Route{
		"GetAllUsers",
		"GET",
		"/all",
		user.GetAllUsers,
	},
	Route{
		"GetUserByID",
		"GET",
		"/userbyid",
		user.GetUserByID,
	},
	Route{
		"GetUserByName",
		"GET",
		"/userbyname",
		user.GetUserByName,
	},
	Route{
		"PostUser",
		"POST",
		"/user",
		user.PostUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/user",
		user.DeleteUser,
	},
}