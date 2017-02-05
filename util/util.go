package util

import (
	"io"
	"io/ioutil"
	"net/http"
)

func GetBodyFromRequest(r *http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	return body
}
