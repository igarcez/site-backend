package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/igarcez/site-backend/app"
	"github.com/igarcez/site-backend/data"
)

func TagIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.Tags{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}
