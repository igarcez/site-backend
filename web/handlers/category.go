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

func CategoryIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.Categories{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}

func CategoryGet(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		var cat app.Category
		db := data.GetConnection()
		result := db.Where("id = ?", id).First(&cat)
		if cat.IsValid() {
			response.Success(w, result)
			return
		}
	}
	err := errors.New("Entity not found")
	response.Fail(w, 404, err, "")
}

func CategoryUpdate(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		body := util.GetBodyFromRequest(r)
		var cat, oldCat app.Category
		if err := json.Unmarshal(body, &cat); err != nil {
			response.Fail(w, 503, err, "")
			return
		}

		id, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			response.Fail(w, 422, err, "id could not be converted to int")
			return
		}

		db := data.GetConnection()
		db.Where("id = ?", id).First(&oldCat)

		cat.ID = uint(id)
		cat.CreatedAt = oldCat.CreatedAt

		if !cat.IsValid() {
			err := errors.New("Changes are not valid")
			response.Fail(w, 503, err, "")
			return
		}
		result := db.Save(&cat)
		response.Success(w, result)
	}
}

func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	body := util.GetBodyFromRequest(r)
	var cat app.Category
	if err := json.Unmarshal(body, &cat); err != nil {
		response.Fail(w, 422, err, "Invalid content")
		return
	}

	if !cat.IsValid() {
		err := errors.New("Entity is not valid")
		response.Fail(w, 400, err, "")
		return
	}

	db := data.GetConnection()
	result := db.Create(&cat)
	response.Success(w, result)
}

func CategoryDelete(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		db := data.GetConnection()
		result := db.Where("id = ?", id).Delete(&app.Category{})
		response.Success(w, result)
	} else {
		err := errors.New("Missing ID argument")
		response.Fail(w, 400, err, "")
	}
}
