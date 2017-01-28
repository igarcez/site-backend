package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igarcez/site-backend/app"
	"github.com/igarcez/site-backend/data"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TypeIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.PageTypes{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}

func CategoryIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.Categories{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}

func PageIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.Pages{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}

func TagIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.Tags{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}
