package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-zoo/bone"
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

func TypeGet(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		var pageType app.PageType
		db := data.GetConnection()
		result := db.Where("id = ?", id).First(&pageType)
		if pageType.IsValid() {
			response.Success(w, result)
			return
		}
	}
	err := errors.New("Entity not found")
	response.Fail(w, 404, err, "")
}

func TypeUpdate(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		body := util.GetBodyFromRequest(r)
		var pageType, oldPageType app.PageType
		if err := json.Unmarshal(body, &pageType); err != nil {
			response.Fail(w, 503, err, "")
			return
		}

		id, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			response.Fail(w, 422, err, "id could not be converted to int")
			return
		}

		db := data.GetConnection()
		db.Where("id = ?", id).First(&oldPageType)

		pageType.ID = uint(id)
		pageType.CreatedAt = oldPageType.CreatedAt

		if !pageType.IsValid() {
			err := errors.New("Changes are not valid")
			response.Fail(w, 503, err, "")
			return
		}
		result := db.Save(&pageType)
		response.Success(w, result)
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

func TypeDelete(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		db := data.GetConnection()
		result := db.Where("id = ?", id).Delete(&app.PageType{})
		response.Success(w, result)
	} else {
		err := errors.New("Missing ID argument")
		response.Fail(w, 400, err, "")
	}
}
