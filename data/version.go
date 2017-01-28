package data

import "time"

type Version struct {
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" sql:"updated_at"`
}
