package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/igarcez/site-backend/app"
	"github.com/igarcez/site-backend/data"
	"github.com/igarcez/site-backend/util"
	"github.com/igarcez/site-backend/web/response"
)

func TypeIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.PageTypes{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}

func TypeCreate(w http.ResponseWriter, r *http.Request) {
	body := util.GetBodyFromRequest(r)
	var pageType app.PageType
	if err := json.Unmarshal(body, &pageType); err != nil {
		response.Fail(w, 422, err, "Invalid content")
		return
	}

	if !pageType.IsValid() {
		err := errors.New("Entity is not valid")
		response.Fail(w, 400, err, "")
		return
	}

	db := data.GetConnection()
	result := db.Create(&pageType)
	response.Success(w, result)
}
