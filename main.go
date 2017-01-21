package main

import "net/http"
import "github.com/igarcez/site-backend/web"

func main() {
	router := web.NewRouter()

	http.ListenAndServe(":8080", router)
}
