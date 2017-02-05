package response

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

func Fail(w http.ResponseWriter, code int, err error, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	response := make(map[string]string)
	response["Error"] = err.Error()
	response["Message"] = message
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func Success(w http.ResponseWriter, response *gorm.DB) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
