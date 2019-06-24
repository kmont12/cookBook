package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"AddRecipe",
		"POST",
		"/recipe/add",
		AddHandler,
	},
	Route{
		"SearchRecipe",
		"GET",
		"/recipe/search",
		SearchHandler,
	},
	Route{
		"UpdateRecipe",
		"POST",
		"/recipe/update",
		SearchHandler,
	},
	Route{
		"AddNotes",
		"POST",
		"/notes",
		AddNotesHandler,
	},
	Route{
		"GetNotes",
		"GET",
		"/notes",
		GetNotesHandler,
	},
	Route{
		"DeleteNotes",
		"DELETE",
		"/notes",
		DeleteNotesHandler,
	},
}
