package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/igarcez/site-backend/app"
	"github.com/igarcez/site-backend/data"
)

func PageIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.Pages{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}
