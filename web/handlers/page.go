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

func PageIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.Pages{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}

func PageGet(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		var page app.Page
		db := data.GetConnection()
		result := db.Where("id = ?", id).First(&page)
		if page.IsValid() {
			response.Success(w, result)
			return
		}
	}
	err := errors.New("Entity not found")
	response.Fail(w, 404, err, "")
}

func PageUpdate(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		body := util.GetBodyFromRequest(r)
		var page, oldPage app.Page
		if err := json.Unmarshal(body, &page); err != nil {
			response.Fail(w, 503, err, "")
			return
		}

		id, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			response.Fail(w, 422, err, "id could not be converted to int")
			return
		}

		db := data.GetConnection()
		db.Where("id = ?", id).First(&oldPage)

		page.ID = uint(id)
		page.CreatedAt = oldPage.CreatedAt

		if !page.IsValid() {
			err := errors.New("Changes are not valid")
			response.Fail(w, 503, err, "")
			return
		}
		result := db.Save(&page)
		response.Success(w, result)
	}
}

func PageCreate(w http.ResponseWriter, r *http.Request) {
	body := util.GetBodyFromRequest(r)
	var page app.Page
	if err := json.Unmarshal(body, &page); err != nil {
		response.Fail(w, 422, err, "Invalid content")
		return
	}

	if !page.IsValid() {
		err := errors.New("Entity is not valid")
		response.Fail(w, 400, err, "")
		return
	}

	db := data.GetConnection()
	result := db.Create(&page)
	response.Success(w, result)
}

func PageDelete(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		db := data.GetConnection()
		result := db.Where("id = ?", id).Delete(&app.Page{})
		response.Success(w, result)
	} else {
		err := errors.New("Missing ID argument")
		response.Fail(w, 400, err, "")
	}
}
