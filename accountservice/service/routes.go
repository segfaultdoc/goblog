package service

import (
	"net/http"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

// Route defines a single route, e.g. a human readable name, HTTP method, pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route


// Initialize our routes
var routes = Routes{
	Route{
		"StoreAccount", // Name
		"POST",         // HTTP method
		"/accounts",    // Route pattern
		StoreAccount,
	},
	Route{
		"UpdateAccount", // Name
		"PUT",           // HTTP method
		"/accounts",     // Route pattern
		UpdateAccount,
	},
	Route{
		"GetAccount",            // Name
		"GET",                   // HTTP method
		"/accounts/{accountId}", // Route pattern
		GetAccount,
	},
	Route{
		"GraphQL",  // Name
		"POST",     // HTTP method
		"/graphql", // Route pattern
		gqlhandler.New(&gqlhandler.Config{
			Schema: &schema,
			Pretty: false,
		}).ServeHTTP,
	},
	Route{
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
	Route{
		"Testability",
		"GET",
		"/testability/healthy/{state}",
		SetHealthyState,
	},
}
