package app

import "github.com/igarcez/site-backend/data"

type PageType struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	data.Version
}

type PageTypes []PageType
