package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route structure to hold the route info
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes array of routes
type Routes []Route

//NewRouter creates a new mux router
func NewRouter() *mux.Router {
	var routes = Routes{
		Route{
			"GetVmsList",
			"GET",
			"/vmlist",
			GetVmsList,
		},
		Route{
			"CreateVm",
			"POST",
			"/createvm",
			CreateVm,
		},
		Route{
			"UpdateVm",
			"PUT",
			"/updatevm",
			UpdateVm,
		},
		Route{
			"DeleteVm",
			"DELETE",
			"/deletevm",
			DeleteVm,
		},
	}

	//Add both the routes with authentication and those without to the router.
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
