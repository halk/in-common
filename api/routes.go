package api

import "net/http"

// Route struct
// Inspired by: http://thenewstack.io/make-a-restful-json-api-go/
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes struct
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AddEvent",
		"POST",
		"/event",
		AddEvent,
	},
	Route{
		"RemoveEvent",
		"DELETE",
		"/event",
		RemoveEvent,
	},
	Route{
		"Recommend",
		"GET",
		"/recommend",
		Recommend,
	},
}
