package main

import (
	"net/http"

	"github.com/igarcez/site-backend/app"
	"github.com/igarcez/site-backend/data"
	"github.com/igarcez/site-backend/web"
)

func main() {
	data.InitDataConnection()
	defer data.CloseDataConnection()

	app.Init()
	router := web.NewRouter()

	http.ListenAndServe(":8080", router)
}
