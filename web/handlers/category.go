package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/igarcez/site-backend/app"
	"github.com/igarcez/site-backend/data"
)

func CategoryIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.Categories{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}
