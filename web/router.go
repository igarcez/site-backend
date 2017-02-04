package web

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/igarcez/site-backend/web/handlers"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type Server struct {
	Routes Routes
}

func NewRouter() *bone.Mux {
	router := bone.New()
	server := InitServer()

	for _, route := range server.Routes {
		router.Register(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}

func (server *Server) add(method string, pattern string, handler http.HandlerFunc) *Server {
	var routeObj = Route{method, pattern, handler}
	server.Routes = append(server.Routes, routeObj)
	return server
}

func InitServer() Server {
	var server Server
	server.add("GET", "/types", handlers.TypeIndex)
	server.add("GET", "/categories", handlers.CategoryIndex)
	server.add("GET", "/pages", handlers.PageIndex)
	server.add("GET", "/tags", handlers.TagIndex)

	server.add("GET", "/type/:id", handlers.TypeGet)
	server.add("GET", "/category/:id", handlers.CategoryGet)
	server.add("GET", "/tag/:id", handlers.TagGet)

	server.add("PUT", "/type/:id", handlers.TypeUpdate)
	server.add("PUT", "/category/:id", handlers.CategoryUpdate)
	server.add("PUT", "/tag/:id", handlers.TagUpdate)

	server.add("POST", "/types", handlers.TypeCreate)
	server.add("POST", "/categories", handlers.CategoryCreate)
	server.add("POST", "/tags", handlers.TagCreate)

	server.add("DELETE", "/type/:id", handlers.TypeDelete)
	server.add("DELETE", "/category/:id", handlers.CategoryDelete)
	server.add("DELETE", "/tag/:id", handlers.TagDelete)
	return server
}
