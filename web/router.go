package web

import (
	"net/http"

	"github.com/go-zoo/bone"
)

type Route struct {
	Name        string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *bone.Mux {
	router := bone.New()

	for _, route := range getRoutes {
		router.GetFunc(route.Pattern, route.HandlerFunc)
	}

	return router
}

var getRoutes = Routes{
	Route{
		"Index",
		"/",
		Index,
	},
	Route{
		"Type Index",
		"/types",
		TypeIndex,
	},
}
