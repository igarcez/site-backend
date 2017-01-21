package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igarcez/site-backend/app"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TypeIndex(w http.ResponseWriter, r *http.Request) {
	types := app.PageTypes{
		app.PageType{
			Code:        "static",
			Description: "static pages",
		},
		app.PageType{
			Code:        "post",
			Description: "Blog style posts",
		},
	}
	if err := json.NewEncoder(w).Encode(types); err != nil {
		panic(err)
	}
}
