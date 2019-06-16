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
		"GET",
		"/recipe/search",
		SearchHandler,
	},
	Route{
		"AddNotes",
		"POST",
		"/notes/add",
		AddNotesHandler,
	},
	Route{
		"GetNotes",
		"POST",
		"/notes/add",
		GetNotesHandler,
	},
}
