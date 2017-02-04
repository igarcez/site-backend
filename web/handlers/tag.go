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

func TagIndex(w http.ResponseWriter, r *http.Request) {
	collection := app.Tags{}
	db := data.GetConnection()
	db.Find(&collection)
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}

func TagGet(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		var tag app.Tag
		db := data.GetConnection()
		result := db.Where("id = ?", id).First(&tag)
		if tag.IsValid() {
			response.Success(w, result)
			return
		}
	}
	err := errors.New("Entity not found")
	response.Fail(w, 404, err, "")
}

func TagUpdate(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		body := util.GetBodyFromRequest(r)
		var tag, oldTag app.Tag
		if err := json.Unmarshal(body, &tag); err != nil {
			response.Fail(w, 503, err, "")
			return
		}

		id, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			response.Fail(w, 422, err, "id could not be converted to int")
			return
		}

		db := data.GetConnection()
		db.Where("id = ?", id).First(&oldTag)

		tag.ID = uint(id)
		tag.CreatedAt = oldTag.CreatedAt

		if !tag.IsValid() {
			err := errors.New("Changes are not valid")
			response.Fail(w, 503, err, "")
			return
		}
		result := db.Save(&tag)
		response.Success(w, result)
	}
}

func TagCreate(w http.ResponseWriter, r *http.Request) {
	body := util.GetBodyFromRequest(r)
	var tag app.Tag
	if err := json.Unmarshal(body, &tag); err != nil {
		response.Fail(w, 422, err, "Invalid content")
		return
	}

	if !tag.IsValid() {
		err := errors.New("Entity is not valid")
		response.Fail(w, 400, err, "")
		return
	}

	db := data.GetConnection()
	result := db.Create(&tag)
	response.Success(w, result)
}

func TagDelete(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")
	if len(id) > 0 {
		db := data.GetConnection()
		result := db.Where("id = ?", id).Delete(&app.Tag{})
		response.Success(w, result)
	} else {
		err := errors.New("Missing ID argument")
		response.Fail(w, 400, err, "")
	}
}
