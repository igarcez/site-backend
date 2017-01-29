package web

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/igarcez/site-backend/web/handlers"
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
		"Type Index",
		"/types",
		handlers.TypeIndex,
	},
	Route{
		"Category Index",
		"/categories",
		handlers.CategoryIndex,
	},
	Route{
		"Page Index",
		"/pages",
		handlers.PageIndex,
	},
	Route{
		"Tag Index",
		"/tags",
		handlers.TagIndex,
	},
}
