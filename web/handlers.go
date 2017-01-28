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
	theType := app.NewPageType()

	collection := theType.GetCollection()

	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}
