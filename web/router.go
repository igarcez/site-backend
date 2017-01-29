package web

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/igarcez/site-backend/web/handlers"
)

type Route struct {
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type Server struct {
	Get    Routes
	Put    Routes
	Delete Routes
	Post   Routes
}

func NewRouter() *bone.Mux {
	router := bone.New()
	server := InitServer()

	for _, route := range server.Get {
		router.GetFunc(route.Pattern, route.HandlerFunc)
	}

	for _, route := range server.Post {
		router.PostFunc(route.Pattern, route.HandlerFunc)
	}

	for _, route := range server.Put {
		router.PutFunc(route.Pattern, route.HandlerFunc)
	}

	for _, route := range server.Delete {
		router.DeleteFunc(route.Pattern, route.HandlerFunc)
	}

	return router
}

func (server *Server) add(verb string, pattern string, handler http.HandlerFunc) *Server {
	var routeObj = Route{pattern, handler}
	var routerList *Routes

	switch verb {
	case "Post":
		routerList = &server.Post
	case "Put":
		routerList = &server.Put
	case "Delete":
		routerList = &server.Delete
	default: // GET is the default
		routerList = &server.Get
	}
	*routerList = append(*routerList, routeObj)
	return server
}

func InitServer() Server {
	var server Server
	// we add the Get verb just for clarity, it is the default
	server.add("Get", "/types", handlers.TypeIndex)
	server.add("Get", "/categories", handlers.CategoryIndex)
	server.add("Get", "/pages", handlers.PageIndex)
	server.add("Get", "/tags", handlers.TagIndex)
	return server
}
