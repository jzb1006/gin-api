package model

import "go-gin-api/util/encode"

type BaseModelField struct {
	CreatedAt     encode.LocalTime `json:"created_at"`
	UpdatedAt     encode.LocalTime `json:"updated_at"`
}
